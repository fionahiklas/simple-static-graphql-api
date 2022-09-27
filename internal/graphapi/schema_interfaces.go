//go:generate mockgen -package graphapi_test -destination=./mock_schema_test.go -source $GOFILE

package graphapi

import (
	"context"
	graphql "github.com/graph-gophers/graphql-go"
)

type AlarmSystemArgs struct {
	AlarmSystemId graphql.ID
}

type HomeQueryArgs struct {
	HomeId graphql.ID
}

type SensorResolver interface {
	Identifier() graphql.ID
	Name() string
	Description() *string
}

type AlarmSystemResolver interface {
	Identifier() graphql.ID
	Name() string
	Description() *string
	Sensors() (*[]*SensorResolver, error)
}

type HomeResolver interface {
	Identifier() graphql.ID
	Name() string
	Description() *string
	AlarmSystem() (*AlarmSystemResolver, error)
}

type RootResolver interface {
	Home(ctx context.Context, args HomeQueryArgs) (*HomeResolver, error)
	Homes(ctx context.Context) (*[]*HomeResolver, error)
	AlarmSystem(ctx context.Context, args AlarmSystemArgs) (*AlarmSystemResolver, error)
	AlarmSystems(ctx context.Context) (*[]*AlarmSystemResolver, error)
}
