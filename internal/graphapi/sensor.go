package graphapi

import (
	"context"

	"github.com/fionahiklas/simple-static-graphql-api/pkg/alarmstorage"

	"github.com/graph-gophers/graphql-go"
)

type SensorResolver struct {
	log    logger
	ctx    context.Context
	sensor *alarmstorage.Sensor
}

func NewSensorResolver(log logger, ctx context.Context, sensor *alarmstorage.Sensor) *SensorResolver {
	return &SensorResolver{
		log:    log,
		ctx:    ctx,
		sensor: sensor,
	}
}

func (sr *SensorResolver) Identifier() graphql.ID {
	return graphql.ID(sr.sensor.Id)
}

func (sr *SensorResolver) Name() string {
	return sr.sensor.Name
}

func (sr *SensorResolver) Description() *string {
	return &sr.sensor.Description
}
