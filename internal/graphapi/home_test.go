package graphapi_test

import (
	"context"
	"github.com/fionahiklas/simple-static-graphql-api/internal/graphapi"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewHome(t *testing.T) {
	result := graphapi.NewHomeResolver(nil, context.Background(), nil)
	require.NotNil(t, result)
}
