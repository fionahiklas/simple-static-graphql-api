//go:generate mockgen -package graphapi_test -destination=./mock_graphapi_test.go -source $GOFILE
package graphapi

import (
	_ "embed"
	"net/http"

	"github.com/fionahiklas/simple-static-graphql-api/pkg/alarmstorage"
	"github.com/graph-gophers/graphql-go/relay"

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
	schema  *graphql.Schema
	storage alarmstorage.ReadAndWrite
}

func NewGraphQLAPI(log logger, storage alarmstorage.ReadAndWrite) *graphQLApi {
	return &graphQLApi{
		schema:  graphql.MustParseSchema(schemaString, NewRoot(log, storage)),
		storage: storage,
	}
}

func (api *graphQLApi) GetHandler() http.Handler {
	return &relay.Handler{
		Schema: api.schema,
	}
}
