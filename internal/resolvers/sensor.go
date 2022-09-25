package resolvers

import (
	"context"
	"github.com/graph-gophers/graphql-go"
)

type SensorResolver struct{}

func NewSensorResolver(ctx context.Context) *SensorResolver {
	return &SensorResolver{}
}

func (sr *SensorResolver) ID() graphql.ID {
	return ""
}

func (sr *SensorResolver) Name() string {
	return ""
}

func (sr *SensorResolver) Description() *string {
	return &emptyString
}
