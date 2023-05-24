//go:build pact

package provider_test

import (
	"fmt"
	"net"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/fionahiklas/simple-static-graphql-api/internal/version"

	"github.com/fionahiklas/simple-static-graphql-api/internal/graphapi"
	"github.com/fionahiklas/simple-static-graphql-api/pkg/alarmstorage"
	"github.com/golang/mock/gomock"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/sirupsen/logrus"
)

func TestProvider_ProviderAllAlarmNames(t *testing.T) {
	const (
		pactConsumerBaseDir = "../../build/pact/consumer"
		pactProviderBaseDir = "../../build/pact/provider"
	)

	// For now don't bother with using a mock for the logger
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	var controller *gomock.Controller
	var mockStorage *MockReadAndWrite

	// Set the initial handler to nil, will change this in reset function
	var handlerWrapper = NewHandlerWrapper(log, nil)

	// NOTE: This is a closure so the value of "t" is captured from the
	// NOTE: enclosing test function.  If you see odd behaviour in the tests
	// NOTE: check this approach!
	resetMocksAndHandler := func() {
		log.Debugf("Resetting mocks and handler")
		controller = gomock.NewController(t)
		mockStorage = NewMockReadAndWrite(controller)
		graphApi := graphapi.NewGraphQLAPI(log, mockStorage)
		graphApiHandler := graphApi.GetHandler()
		handlerWrapper.ChangeHandler(graphApiHandler)
	}

	port := startHTTPServerForHandler(log, handlerWrapper)

	pactInstance := &dsl.Pact{
		Provider: "apiprovider",
		Consumer: "apiconsumer",
		PactDir:  pactProviderBaseDir,
		// This doesn't appear to be honoured for VerifyProvider
		LogDir: pactProviderBaseDir + "/logs",
	}

	providerBaseUrl := fmt.Sprintf("http://localhost:%d", port)
	log.Debugf("Provider host: %s", providerBaseUrl)

	_, err := pactInstance.VerifyProvider(t, types.VerifyRequest{
		// Doesn't seem to honour the LogDir in the dsl.Pact instance, need to set here
		PactLogDir:      pactProviderBaseDir + "/logs",
		Provider:        pactInstance.Provider,
		ProviderBaseURL: providerBaseUrl,
		PactURLs:        []string{filepath.ToSlash(fmt.Sprintf("%s/apiconsumer-apiprovider.json", pactConsumerBaseDir))},
		// This is for PactFlow features
		//BrokerURL:                  "https://<org domain or name>.pactflow.io",
		//BrokerToken:                os.Getenv("PACT_BROKER_TOKEN"),
		ProviderVersion:            version.CodeVersion(),
		PublishVerificationResults: true,
		StateHandlers: types.StateHandlers{
			// Setup any state required by the test
			// in this case, we ensure there is a "user" in the system
			"Two Alarms Exist": func() error {
				log.Debugf("Setting up 'Two Alarms Exist' state")
				resetMocksAndHandler()
				mockStorage.EXPECT().GetAlarms().Return([]*alarmstorage.Alarm{
					&alarmstorage.Alarm{
						Identifier:  "1",
						Name:        "Greebo",
						Description: "",
						Sensors:     nil,
					},
					&alarmstorage.Alarm{
						Identifier:  "2",
						Name:        "Nanny Ogg",
						Description: "",
						Sensors:     nil,
					},
				})
				return nil
			},

			"One Alarm exists": func() error {
				log.Debugf("Setting up 'One Alarm exists' state")
				resetMocksAndHandler()
				mockStorage.EXPECT().GetAlarm("Cat_01").Return(
					&alarmstorage.Alarm{
						Identifier:  "Cat_01",
						Name:        "Greebo",
						Description: "Cat",
						Sensors:     nil,
					})
				return nil
			},
		},
	})

	if err != nil {
		log.Debugf("Error in pact provider %s", err)
	}

}

func startHTTPServerForHandler(log *logrus.Logger, handlerWrapper *handlerWrapper) int {
	mux := http.NewServeMux()
	mux.Handle("/graphql", handlerWrapper)

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	port := listener.Addr().(*net.TCPAddr).Port
	log.Debugf("Using port: %d", port)

	go http.Serve(listener, mux)
	return port
}

type handlerWrapper struct {
	log           *logrus.Logger
	handlerToWrap http.Handler
}

func NewHandlerWrapper(log *logrus.Logger, handler http.Handler) *handlerWrapper {
	return &handlerWrapper{
		log:           log,
		handlerToWrap: handler,
	}
}

func (hw *handlerWrapper) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	hw.log.Debugf("Handler wrapper, request, url: %s, method: %s", request.URL.String(), request.Method)
	hw.handlerToWrap.ServeHTTP(response, request)
}

func (hw *handlerWrapper) ChangeHandler(handler http.Handler) {
	hw.log.Debugf("Switching handlers")
	hw.handlerToWrap = handler
}
