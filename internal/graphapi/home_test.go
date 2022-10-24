package graphapi_test

import (
	"context"
	"testing"

	"github.com/fionahiklas/simple-static-graphql-api/internal/graphapi"

	"github.com/stretchr/testify/require"
)

func TestNewHome(t *testing.T) {
	result := graphapi.NewHomeResolver(nil, context.Background(), nil)
	require.NotNil(t, result)
}
