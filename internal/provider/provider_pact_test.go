//go:build pact

package provider_test

import (
	"github.com/golang/mock/gomock"
	"github.com/pact-foundation/pact-go/dsl"
	"testing"
)

func TestProvider_PactTests(t *testing.T) {

	const (
		pactBaseDir = "../../build/pact/provider"
	)
	// For now don't bother with using a mock for the logger
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	pact := &dsl.Pact{
		Provider: "apiprovider",
		PactDir:  pactBaseDir,
		LogDir:   pactBaseDir + "/logs",
	}

	ctrl := gomock.NewController(t)
	mockStorage := NewMockReadAndWrite(ctrl)

}
