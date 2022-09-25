package graphapi_test

import (
	"github.com/fionahiklas/simple-static-graphql-api/internal/graphapi"
	"github.com/fionahiklas/simple-static-graphql-api/internal/resolvers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSchema(t *testing.T) {

	t.Run("schema and resolvers match", func(t *testing.T) {
		result := graphapi.NewGraphQL(resolvers.NewRoot())
		require.NotNil(t, result)
	})
}
