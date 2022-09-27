package resolvers

import (
	"context"
	"errors"

	"github.com/graph-gophers/graphql-go"
)

type AlarmSystemResolver struct{}

func NewAlarmSystemResolver(ctx context.Context) *AlarmSystemResolver {
	return &AlarmSystemResolver{}
}

func (ar *AlarmSystemResolver) Identifier() graphql.ID {
	return ""
}

func (ar *AlarmSystemResolver) Name() string {
	return ""
}

func (ar *AlarmSystemResolver) Description() *string {
	return &emptyString
}

func (ar *AlarmSystemResolver) Sensors() (*[]*SensorResolver, error) {
	return nil, errors.New("not implemented")
}
