package resolvers_test

import (
	"context"
	"testing"

	"github.com/fionahiklas/simple-static-graphql-api/internal/resolvers"
	"github.com/stretchr/testify/require"
)

func TestNewHome(t *testing.T) {
	result := resolvers.NewHomeResolver(context.Background())
	require.NotNil(t, result)
}
