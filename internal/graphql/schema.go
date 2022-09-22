package graphql

import (
	_ "embed"
	"net/http"
)

//go:embed schema.graphql
var schemaString string

func NewGraphQLAPI() http.Handler {
	return nil
}
