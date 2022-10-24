package consumer_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"

	"github.com/fionahiklas/simple-static-graphql-api/internal/consumer"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestConsumer_GetAllAlarmNames(t *testing.T) {
	const (
		testCallUrl       = "https://ankh.morpork.discworld"
		testHappyResponse = `{"data":{"alarmSystems":[{"name":"Alarm One"},{"name":"Alarm Two"},{"name":"Alarm Three"}]}}`
	)

	// For now don't bother with using a mock for the logger
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	var mockHttpClient *MockhttpClient

	resetMocks := func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockHttpClient = NewMockhttpClient(ctrl)
	}

	buildResponse := func(status int, body string) *http.Response {
		return &http.Response{
			StatusCode: status,
			Body:       io.NopCloser(strings.NewReader(body)),
		}
	}

	t.Run("happy path", func(t *testing.T) {
		resetMocks(t)
		consumerToTest := consumer.NewConsumer(log, mockHttpClient, testCallUrl)

		mockHttpClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(request *http.Request) (*http.Response, error) {
			require.Equal(t, testCallUrl, request.URL.String())
			return buildResponse(http.StatusOK, testHappyResponse), nil
		})

		alarmResult, err := consumerToTest.GetAllAlarmNames()
		require.NoError(t, err)
		require.Equal(t, 3, len(alarmResult))
	})

	t.Run("pact consumer test", func(t *testing.T) {
		pactInstance := dsl.Pact{
			// TODO: Are there names significant?
			Consumer: "apiconsumer",
			Provider: "apiprovider",
		}

		// The 'true' argument means "start the mock server"
		pactInstance.Setup(true)

		log.Debugf("Pact server host: %s", pactInstance.Host)
		log.Debugf("Pact server port: %d", pactInstance.Server.Port)
	})
}
