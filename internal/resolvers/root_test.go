package resolvers_test

import (
	"github.com/fionahiklas/simple-static-graphql-api/internal/resolvers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewRoot(t *testing.T) {
	result := resolvers.NewRoot()
	require.NotNil(t, result)
}
