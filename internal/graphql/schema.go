package graphql

import (
	_ "embed"
	graphqlgo "github.com/graph-gophers/graphql-go"
	"net/http"
)

//go:embed schema.graphql
var schemaString string

type graphQLApi struct {
	schema *graphqlgo.Schema
}

func NewGraphQL(rootResolver interface{}) *graphQLApi {
	return &graphQLApi{
		schema: graphqlgo.MustParseSchema(schemaString, rootResolver),
	}
}

func (g *graphQLApi) ServerHTTP(w http.ResponseWriter, r *http.Request) {

}
