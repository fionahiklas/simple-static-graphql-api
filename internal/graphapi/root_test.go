package graphapi_test

import (
	"github.com/fionahiklas/simple-static-graphql-api/internal/graphapi"
	"github.com/sirupsen/logrus"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRoot(t *testing.T) {
	log := logrus.New()
	result := graphapi.NewRoot(log)
	require.NotNil(t, result)
}
