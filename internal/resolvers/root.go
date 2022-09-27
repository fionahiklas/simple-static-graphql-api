package resolvers

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

type queryResolver struct{}

func NewRoot() *queryResolver {
	return &queryResolver{}
}

func (qr *queryResolver) Homes(ctx context.Context) (*[]*HomeResolver, error) {
	return nil, nil
}

type HomeQueryArgs struct {
	HomeId graphql.ID
}

func (qr *queryResolver) Home(ctx context.Context, args HomeQueryArgs) (*HomeResolver, error) {
	return nil, nil
}

type AlarmSystemArgs struct {
	AlarmSystemId graphql.ID
}

func (qr *queryResolver) AlarmSystem(ctx context.Context, args AlarmSystemArgs) (*AlarmSystemResolver, error) {
	return nil, nil
}

func (qr *queryResolver) AlarmSystems(ctx context.Context) (*[]*AlarmSystemResolver, error) {
	return nil, nil
}
