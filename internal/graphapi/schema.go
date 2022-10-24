package graphapi

import (
	_ "embed"
	"github.com/graph-gophers/graphql-go/relay"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
)

type logger interface {
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

//go:embed schema.graphql
var schemaString string

type graphQLApi struct {
	schema *graphql.Schema
}

func NewGraphQLAPI(log logger) *graphQLApi {
	return &graphQLApi{
		schema: graphql.MustParseSchema(schemaString, NewRoot(log)),
	}
}

func (api *graphQLApi) GetHandler() http.Handler {
	return &relay.Handler{
		Schema: api.schema,
	}
}
