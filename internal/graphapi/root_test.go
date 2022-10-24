package graphapi_test

import (
	"testing"

	"github.com/fionahiklas/simple-static-graphql-api/internal/graphapi"
	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/require"
)

func TestNewRoot(t *testing.T) {
	log := logrus.New()
	result := graphapi.NewRoot(log, nil)
	require.NotNil(t, result)
}
