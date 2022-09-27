package resolvers_test

import (
	"testing"

	"github.com/fionahiklas/simple-static-graphql-api/internal/resolvers"
	"github.com/stretchr/testify/require"
)

func TestNewRoot(t *testing.T) {
	result := resolvers.NewRoot()
	require.NotNil(t, result)
}
