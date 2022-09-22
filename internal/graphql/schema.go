package graphql

import (
	_ "embed"
	graphqlgo "github.com/graph-gophers/graphql-go"
	"net/http"
)

//go:embed schema.graphql
var schemaString string

func NewGraphQLAPI() http.Handler {
	graphqlgo.MustParseSchema(schemaString, nil)
	return nil
}
