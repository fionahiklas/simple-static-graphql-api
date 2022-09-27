package resolvers

import (
	"context"
	"errors"

	"github.com/graph-gophers/graphql-go"
)

var emptyString = ""

type HomeResolver struct{}

func NewHomeResolver(ctx context.Context) *HomeResolver {
	return &HomeResolver{}
}

func (hr *HomeResolver) ID() graphql.ID {
	return ""
}

func (hr *HomeResolver) Name() string {
	return ""
}

func (hr *HomeResolver) Description() *string {
	return &emptyString
}

func (hr *HomeResolver) AlarmSystem() (*AlarmSystemResolver, error) {
	return nil, errors.New("not implemented")
}
