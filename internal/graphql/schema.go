package graphql

import (
	_ "embed"
	graphqlgo "github.com/graph-gophers/graphql-go"
	"net/http"
)

//go:embed schema.graphql
var schemaString string

type graphQLApi struct{}

func NewGraphQL() *graphQLApi {
	graphqlgo.MustParseSchema(schemaString, nil)
	return nil
}

func (g *graphQLApi) ServerHTTP(w http.ResponseWriter, r *http.Request) {

}
