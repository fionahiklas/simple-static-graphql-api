//go:build pact

package consumer_test

import (
	"testing"

	"github.com/pact-foundation/pact-go/dsl"

	"github.com/sirupsen/logrus"
)

func TestConsumer_GetAllAlarmNames(t *testing.T) {
	const (
		testCallUrl       = "https://ankh.morpork.discworld"
		testHappyResponse = `{"data":{"alarmSystems":[{"name":"Alarm One"},{"name":"Alarm Two"},{"name":"Alarm Three"}]}}`
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

		log.Debugf("Pact server host: %s", pactInstance.Host)
		log.Debugf("Pact server port: %d", pactInstance.Server.Port)
	})
}
