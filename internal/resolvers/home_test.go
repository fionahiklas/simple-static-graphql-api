package resolvers_test

import (
	"context"
	"github.com/fionahiklas/simple-static-graphql-api/internal/resolvers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewHome(t *testing.T) {
	result := resolvers.NewHomeResolver(context.Background())
	require.NotNil(t, result)
}
