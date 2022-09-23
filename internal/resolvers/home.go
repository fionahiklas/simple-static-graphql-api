package resolvers

import (
	"context"
	graphqlgo "github.com/graph-gophers/graphql-go"
)

type HomeResolver struct{}

func NewHome(ctx context.Context) *HomeResolver {
	return nil
}

func (hr *HomeResolver) ID() graphqlgo.ID {
	return ""
}

func (hr *HomeResolver) Name() string {
	return ""
}

func (hr *HomeResolver) Description() string {
	return ""
}

func (hr *HomeResolver) AlarmSystem() string {
	return ""
}
