package graphapi

import (
	"context"
	"errors"

	"github.com/fionahiklas/simple-static-graphql-api/pkg/alarmstorage"

	"github.com/graph-gophers/graphql-go"
)

type HomeResolver struct {
	log  logger
	ctx  context.Context
	home *alarmstorage.Home
}

func NewHomeResolver(log logger, ctx context.Context, home *alarmstorage.Home) *HomeResolver {
	return &HomeResolver{
		log:  log,
		ctx:  ctx,
		home: home,
	}
}

func (hr *HomeResolver) Identifier() graphql.ID {
	return graphql.ID(hr.home.Id)
}

func (hr *HomeResolver) Name() string {
	return hr.home.Name
}

func (hr *HomeResolver) Description() *string {
	return &hr.home.Description
}

func (hr *HomeResolver) AlarmSystem() (*AlarmSystemResolver, error) {
	return nil, errors.New("not implemented")
}
