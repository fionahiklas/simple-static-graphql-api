package graphapi

import (
	"context"

	"github.com/fionahiklas/simple-static-graphql-api/pkg/alarmstorage"

	"github.com/graph-gophers/graphql-go"
)

type queryResolver struct {
	log     logger
	storage alarmstorage.ReadAndWrite
}

func NewRoot(log logger, storage alarmstorage.ReadAndWrite) *queryResolver {
	return &queryResolver{
		log:     log,
		storage: storage,
	}
}

func (qr *queryResolver) Homes(ctx context.Context) (*[]*HomeResolver, error) {
	homesFromStorage := qr.storage.GetHomes()
	var homesResult []*HomeResolver = nil
	if homesFromStorage != nil {
		homesResult = make([]*HomeResolver, 0, len(homesFromStorage))
		for _, home := range homesFromStorage {
			homesResult = append(homesResult, NewHomeResolver(qr.log, ctx, home))
		}
	}
	return &homesResult, nil
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
	alarmsFromStorage := qr.storage.GetAlarms()
	var alarmsResult []*AlarmSystemResolver = nil
	if alarmsFromStorage != nil {
		alarmsResult = make([]*AlarmSystemResolver, 0, len(alarmsFromStorage))
		for _, alarm := range alarmsFromStorage {
			alarmsResult = append(alarmsResult, NewAlarmSystemResolver(qr.log, ctx, alarm))
		}
	}
	return &alarmsResult, nil

}
