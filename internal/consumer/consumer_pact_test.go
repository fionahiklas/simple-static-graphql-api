//go:build pact

package consumer_test

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/fionahiklas/simple-static-graphql-api/internal/consumer"

	"github.com/pact-foundation/pact-go/dsl"

	"github.com/sirupsen/logrus"
)

// Note that we *might* want to name the tests differently but this is actually
// fine because of the build constraint at the top of this file.  In short this
// code and the regular test code will never be compiled/linked together
func TestConsumer_GetAllAlarmNames(t *testing.T) {
	const (
		testHappyResponse = `{"data":{"alarmSystems":[{"name":"Greebo"},{"name":"Nanny Ogg"}]}}`
		graphQLPathSuffix = "/graphql"
	)

	// For now don't bother with using a mock for the logger
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	t.Run("pact consumer test", func(t *testing.T) {
		pactInstance := dsl.Pact{
			// TODO: Are there names significant?
			Consumer: "apiconsumer",
			Provider: "apiprovider",
		}

		// The 'true' argument means "start the mock server"
		pactInstance.Setup(true)

		// This will stop the mock server
		defer pactInstance.Teardown()

		log.Debugf("Pact server host: %s", pactInstance.Host)
		log.Debugf("Pact server port: %d", pactInstance.Server.Port)

		pactInstance.AddInteraction().
			Given("Two Alarms Exist").
			UponReceiving("All Alarms is requested").
			WithRequest(dsl.Request{
				Method: http.MethodPost,
				Path:   dsl.Term("/graphql"),
				// Using trim to make the body slightly different to the one the consumer really sends
				Body:    strings.TrimSpace(consumer.AllAlarmSystemQuery),
				Headers: http.Header{"Content-Type": {"application/json"}},
			}).
			WillRespondWith(dsl.Response{
				Status:  http.StatusOK,
				Body:    testHappyResponse,
				Headers: http.Header{"Content-Type": {"application/json"}},
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

		// If all is well then write the pact test
		if err := pactInstance.WritePact(); err != nil {
			t.Fatalf("Failed to write the pact test: %s", err)
		}
	})
}
