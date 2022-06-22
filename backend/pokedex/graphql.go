package pokedex

import (
	"context"
	_ "embed"
	"fmt"
	"sort"
	"strings"

	"flamingo.me/graphql"
	"github.com/bastianccm/list"
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
type Resolver struct {
	baseURL string
}

func (resolver *Resolver) Inject(cfg *struct {
	BaseURL string `inject:"config:pokedex.baseurl"`
}) {
	resolver.baseURL = strings.TrimRight(cfg.BaseURL, "/")
}

func (resolver *Resolver) Pokemon(ctx context.Context, ids []int, catched *bool) ([]*Pokemon, error) {
	offset := 0
	limit := 151
	if len(ids) > 0 {
		sort.Ints(ids)
		ids = list.Filter(ids, func(item int) bool { return item <= 151 })
		if len(ids) > 0 {
			offset = ids[0] - 1
			limit = ids[len(ids)-1]
		}
	}

	pokemon, err := fetchPokemonList(ctx, fmt.Sprintf("%s/pokemon?offset=%d&limit=%d", resolver.baseURL, offset, limit))
	if err != nil {
		return nil, fmt.Errorf("unable to fetch pokemon: %w", err)
	}

	if len(ids) > 0 {
		pokemon = list.Filter(pokemon, func(pokemon *Pokemon) bool { return list.Contains(ids, pokemon.ID) })
	}
	if catched != nil {
		pokemon = list.Filter(pokemon, func(p *Pokemon) bool { return p.Catched == *catched })
	}

	return pokemon, nil
}

func (resolver *Resolver) Total(ctx context.Context) (int, error) {
	return 151, nil
}

func (resolver *Resolver) TotalCatched(ctx context.Context) (int, error) {
	return 0, nil
}

func (resolver *Resolver) SetCatched(ctx context.Context, id int, catched bool) (*Pokemon, error) {
	return nil, nil
}
