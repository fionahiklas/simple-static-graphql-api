package graphapi_test

import (
	"github.com/sirupsen/logrus"
	"testing"

	"github.com/fionahiklas/simple-static-graphql-api/internal/graphapi"
	"github.com/stretchr/testify/require"
)

func TestSchema(t *testing.T) {

	t.Run("schema and resolvers match", func(t *testing.T) {
		log := logrus.New()
		result := graphapi.NewGraphQLAPI(log, nil)
		require.NotNil(t, result)
	})
}
