//go:build pact

package consumer_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/fionahiklas/simple-static-graphql-api/internal/consumer"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

type JSONNode = map[string]interface{}

func TestConsumer_PactContract(t *testing.T) {
	const (
		pactBaseDir            = "../../build/pact/consumer"
		testHappyResponse      = `{"data":{"alarmSystems":[{"name":"Greebo"},{"name":"Nanny Ogg"}]}}`
		oneAlarmSystemResponse = `{"data":{"alarmSystem": { "name": "Greebo", "identifier": "Cat_01", "description": "Cat" }}}`
		graphQLPathSuffix      = "/graphql"
		consumerName           = "apiconsumer"
		providerName           = "apiprovider"
	)

	// For now don't bother with using a mock for the logger
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	pactInstance := dsl.Pact{
		Consumer: "apiconsumer",
		Provider: "apiprovider",
		LogDir:   pactBaseDir + "/logs",
		PactDir:  pactBaseDir,
	}

	// The 'true' argument means "start the mock server"
	pactInstance.Setup(true)

	// This will stop the mock server
	defer pactInstance.Teardown()

	log.Debugf("Pact server host: %s", pactInstance.Host)
	log.Debugf("Pact server port: %d", pactInstance.Server.Port)

	urlPath := fmt.Sprintf("http://%s:%d%s", pactInstance.Host, pactInstance.Server.Port, graphQLPathSuffix)
	log.Debugf("Using this for the URL: %s", urlPath)

	t.Run("all alarm names", func(t *testing.T) {

		pactInstance.AddInteraction().
			Given("Two Alarms Exist").
			UponReceiving("All Alarms is requested").
			WithRequest(dsl.Request{
				Method: http.MethodPost,
				Path:   dsl.String(graphQLPathSuffix),
				// Using trim to make the body slightly different to the one the consumer really sends
				Body: JSONNode{
					"query": "{ alarmSystems { name } }",
				},
				// Not sure how this matcher works with multiple values for headers, didn't like an array
				Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			}).
			WillRespondWith(dsl.Response{
				Status: http.StatusOK,
				Body:   convertStringToJSONNodes(t, testHappyResponse),
				// Not sure how this matcher works with multiple values for headers, didn't like an array
				Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			})

		err := pactInstance.Verify(func() error {
			consumerToTest := consumer.NewConsumer(log, &http.Client{}, urlPath)

			result, err := consumerToTest.GetAllAlarmNames()

			require.NoError(t, err, "Error from GetAllAlarmNames")
			require.Equal(t, 2, len(result))

			return err
		})

		if err != nil {
			t.Fatalf("Pact test verify failed with error: %s", err)
		}
	})

	t.Run("get one alarm", func(t *testing.T) {
		pactInstance.AddInteraction().
			Given("One Alarm exists").
			UponReceiving("Only one alarm is requested").
			WithRequest(dsl.Request{
				Method: http.MethodPost,
				Path:   dsl.String(graphQLPathSuffix),
				// Using trim to make the body slightly different to the one the consumer really sends
				Body: JSONNode{
					"query": "query($identifier: ID!){ alarmSystem(alarmSystemId: $identifier){ identifier name description  } }",
					"variables": JSONNode{
						"identifier": "Cat_01",
					},
				},
				// Not sure how this matcher works with multiple values for headers, didn't like an array
				Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			}).
			WillRespondWith(dsl.Response{
				Status: http.StatusOK,
				Body:   convertStringToJSONNodes(t, oneAlarmSystemResponse),
				// Not sure how this matcher works with multiple values for headers, didn't like an array
				Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			})

		err := pactInstance.Verify(func() error {
			consumerToTest := consumer.NewConsumer(log, &http.Client{}, urlPath)

			oneAlarm, err := consumerToTest.GetOneAlarm("Cat_01")
			require.NoError(t, err)
			require.NotNil(t, oneAlarm)
			require.Equal(t, "Cat_01", oneAlarm.Identifier)
			require.Equal(t, "Greebo", oneAlarm.Name)
			require.Equal(t, "Cat", oneAlarm.Description)

			return err
		})

		if err != nil {
			t.Fatalf("Pact test verify failed with error: %s", err)
		}
	})

	// If all is well then write the pact test
	if err := pactInstance.WritePact(); err != nil {
		t.Fatalf("Failed to write the pact test: %s", err)
	}

	// This is only needed for use with PactFlow publishing
	//
	//publisher := dsl.Publisher{}
	//publisher.Publish(types.PublishRequest{
	//	PactURLs:        []string{pactBaseDir + "/" + consumerName + "-" + providerName + ".json"},
	//	PactBroker:      "https://<org domain or name>.pactflow.io",
	//	BrokerToken:     os.Getenv("PACT_BROKER_TOKEN"),
	//	ConsumerVersion: version.CodeVersion(),
	//})
}

func convertStringToJSONNodes(t *testing.T, textToMarshal string) JSONNode {
	var jsonNodeData JSONNode
	err := json.Unmarshal([]byte(textToMarshal), &jsonNodeData)
	if err != nil {
		t.Fatal(err)
	}
	return jsonNodeData
}
