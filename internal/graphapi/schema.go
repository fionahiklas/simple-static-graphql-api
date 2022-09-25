package graphapi

import (
	_ "embed"
	graphql "github.com/graph-gophers/graphql-go"
	"net/http"
)

//go:embed schema.graphql
var schemaString string

type graphQLApi struct {
	schema *graphql.Schema
}

func NewGraphQL(rootResolver interface{}) *graphQLApi {
	return &graphQLApi{
		schema: graphql.MustParseSchema(schemaString, rootResolver),
	}
}

func (g *graphQLApi) ServerHTTP(w http.ResponseWriter, r *http.Request) {

}
