//go:build pact

package consumer_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/fionahiklas/simple-static-graphql-api/internal/consumer"

	"github.com/pact-foundation/pact-go/dsl"

	"github.com/sirupsen/logrus"
)

type JSONNode = map[string]interface{}

// Note that we *might* want to name the tests differently but this is actually
// fine because of the build constraint at the top of this file.  In short this
// code and the regular test code will never be compiled/linked together
func TestConsumer_GetAllAlarmNames(t *testing.T) {
	const (
		pactConsumerBaseDir    = "../../build/pact/consumer"
		testAllAlarmsNameQuery = `{"query": "{ alarmSystems { name } }" }`
		testHappyResponse      = `{"data":{"alarmSystems":[{"name":"Greebo"},{"name":"Nanny Ogg"}]}}`
		graphQLPathSuffix      = "/graphql"
	)

	// For now don't bother with using a mock for the logger
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	pactInstance := dsl.Pact{
		// These names are used to identify provider/consumer in Pact test output
		Consumer: "apiconsumer",
		Provider: "apiprovider",
		LogDir:   pactConsumerBaseDir + "/logs",
		PactDir:  pactConsumerBaseDir,
	}

	// The 'true' argument means "start the mock server"
	pactInstance.Setup(true)

	// This will stop the mock server
	defer pactInstance.Teardown()

	log.Debugf("Pact server host: %s", pactInstance.Host)
	log.Debugf("Pact server port: %d", pactInstance.Server.Port)

	t.Run("pact consumer test", func(t *testing.T) {

		pactInstance.AddInteraction().
			Given("Two Alarms Exist").
			UponReceiving("All Alarms is requested").
			WithRequest(dsl.Request{
				Method: http.MethodPost,
				Path:   dsl.String(graphQLPathSuffix),
				// Using trim to make the body slightly different to the one the consumer really sends
				Body: convertJSONStringToJSONNodes(t, testAllAlarmsNameQuery),
				// Not sure how this matcher works with multiple values for headers, didn't like an array
				Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			}).
			WillRespondWith(dsl.Response{
				Status: http.StatusOK,
				Body:   convertJSONStringToJSONNodes(t, testHappyResponse),
				// Not sure how this matcher works with multiple values for headers, didn't like an array
				Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			})

		err := pactInstance.Verify(func() error {
			urlPath := fmt.Sprintf("http://%s:%d%s", pactInstance.Host, pactInstance.Server.Port, graphQLPathSuffix)
			log.Debugf("Using this for the URL: %s", urlPath)
			consumerToTest := consumer.NewConsumer(log, &http.Client{}, urlPath)

			result, err := consumerToTest.GetAllAlarmNames()

			if err != nil {
				log.Debugf("Error from GetAllAlarmNames: %s", err)
				return err
			}

			if len(result) != 2 {
				log.Debugf("Didn't get 2 results, got: %d", len(result))
				return errors.New("Didn't get enough data back")
			}

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
}

func convertJSONStringToJSONNodes(t *testing.T, textToMarshal string) JSONNode {
	var jsonNodeData JSONNode
	err := json.Unmarshal([]byte(textToMarshal), &jsonNodeData)
	if err != nil {
		t.Fatal(err)
	}
	return jsonNodeData
}
