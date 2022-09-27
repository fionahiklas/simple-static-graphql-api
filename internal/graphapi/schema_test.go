package graphapi_test

import (
	"github.com/golang/mock/gomock"
	"testing"

	"github.com/fionahiklas/simple-static-graphql-api/internal/graphapi"
	"github.com/fionahiklas/simple-static-graphql-api/internal/resolvers"
	"github.com/stretchr/testify/require"
)

func TestSchema(t *testing.T) {

	t.Run("schema and resolvers match", func(t *testing.T) {
		result := graphapi.NewGraphQL(resolvers.NewRoot())
		require.NotNil(t, result)
	})

	t.Run("schema and interfaces match", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		rootResolverMock := NewMockRootResolver(ctrl)
		result := graphapi.NewGraphQL(rootResolverMock)
		require.NotNil(t, result)
	})
}
