package pokedex

import (
	"context"
	_ "embed"

	"flamingo.me/graphql"
)

type graphqlService struct{}

//go:embed schema.graphql
var schema []byte

// Schema getter
func (*graphqlService) Schema() []byte {
	return schema
}

// Types mapping of GraphQL types to Go
func (*graphqlService) Types(types *graphql.Types) {
	types.Resolve("Query", "pokemon", Resolver{}, "Pokemon")
	types.Resolve("Query", "total", Resolver{}, "Total")
	types.Resolve("Query", "totalCatched", Resolver{}, "TotalCatched")
	types.Map("Pokemon", new(Pokemon))

	types.Resolve("Mutation", "setCatched", Resolver{}, "SetCatched")
}

// Resolver definition
type Resolver struct{}

// Inject dependencies
func (resolver *Resolver) Inject() {
	// resolver.service = service
}

func (resolver *Resolver) Pokemon(ctx context.Context, ids []string, catched *bool) (pokemon []*Pokemon, err error) {
	return nil, nil
}

func (resolver *Resolver) Total(ctx context.Context) (int, error) {
	return 10, nil
}

func (resolver *Resolver) TotalCatched(ctx context.Context) (int, error) {
	return 10, nil
}

func (resolver *Resolver) SetCatched(ctx context.Context, id string, catched bool) (*Pokemon, error) {
	return nil, nil
}
