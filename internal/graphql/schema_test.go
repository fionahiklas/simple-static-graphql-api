package graphql_test

import (
	"github.com/fionahiklas/simple-static-graphql-api/internal/graphql"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSchema(t *testing.T) {

	t.Run("embedded string returned", func(t *testing.T) {
		result := graphql.NewGraphQL(nil)
		require.NotNil(t, result)
	})
}
