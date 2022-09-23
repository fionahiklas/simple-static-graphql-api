package resolvers

import "context"

type queryResolver struct{}

func NewRoot() *queryResolver {
	return &queryResolver{}
}

func (qr *queryResolver) Homes(ctx context.Context) (*[]*HomeResolver, error) {
	return nil, nil
}
