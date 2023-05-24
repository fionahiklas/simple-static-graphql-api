package graphapi

import (
	"context"
	"errors"

	"github.com/fionahiklas/simple-static-graphql-api/pkg/alarmstorage"

	"github.com/graph-gophers/graphql-go"
)

type AlarmSystemResolver struct {
	log   logger
	ctx   context.Context
	alarm *alarmstorage.Alarm
}

func NewAlarmSystemResolver(log logger, ctx context.Context, alarm *alarmstorage.Alarm) *AlarmSystemResolver {
	return &AlarmSystemResolver{
		log:   log,
		ctx:   ctx,
		alarm: alarm,
	}
}

func (ar *AlarmSystemResolver) Identifier() graphql.ID {
	return graphql.ID(ar.alarm.Identifier)
}

func (ar *AlarmSystemResolver) Name() string {
	return ar.alarm.Name
}

func (ar *AlarmSystemResolver) Description() *string {
	return &ar.alarm.Description
}

func (ar *AlarmSystemResolver) Sensors() (*[]*SensorResolver, error) {
	return nil, errors.New("not implemented")
}
