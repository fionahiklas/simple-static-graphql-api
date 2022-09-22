package resolvers

type queryResolver struct{}

func NewRoot() *queryResolver {
	return &queryResolver{}
}
