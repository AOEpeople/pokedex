package pokedex

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bastianccm/list"
	"golang.org/x/sync/errgroup"
)

type Pokemon struct {
	ID      int
	Name    string
	Type    []string
	Catched bool
}

func fetchPokemonList(ctx context.Context, url string) ([]*Pokemon, error) {
	type pokemonListEntryDto struct {
		Name string
		URL  string
	}
	type pokemonListDto struct {
		Results []pokemonListEntryDto
	}

	pokemonList, err := fetch[pokemonListDto](ctx, url)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch pokemon %s: %w", url, err)
	}

	errGroup, ctx := errgroup.WithContext(ctx)
	pokemonChannels := list.Map(pokemonList.Results, func(pokemon pokemonListEntryDto) <-chan *Pokemon {
		result := make(chan *Pokemon, 1)
		errGroup.Go(func() error {
			pokemon, err := fetchPokemon(ctx, pokemon.URL)
			if err != nil {
				return err
			}
			result <- pokemon
			return nil
		})
		return result
	})

	if err := errGroup.Wait(); err != nil {
		return nil, fmt.Errorf("error getting pokemon: %w", err)
	}

	pokemons := list.Map(pokemonChannels, func(pokemonChannel <-chan *Pokemon) *Pokemon { return <-pokemonChannel })
	return list.Sort(pokemons, func(l, r *Pokemon) bool { return l.ID < r.ID }), nil
}

func initCatched() map[int]bool {
	return make(map[int]bool, 151)
}

var catched = initCatched()

func setCatched(id int) {
	if id <= 0 || id > 151 {
		return
	}
	catched[id] = true
}

func unsetCatched(id int) {
	delete(catched, id)
}

func fetchPokemon(ctx context.Context, url string) (*Pokemon, error) {
	type pokemonTypeDto struct {
		Type struct {
			Name string
		}
	}

	type pokemonDto struct {
		ID    int
		Name  string
		Types []pokemonTypeDto
	}

	pokemon, err := fetch[pokemonDto](ctx, url)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch pokemon %s: %w", url, err)
	}

	return &Pokemon{
		ID:   pokemon.ID,
		Name: pokemon.Name,
		Type: list.Map(pokemon.Types, func(typ pokemonTypeDto) string {
			return typ.Type.Name
		}),
		Catched: catched[pokemon.ID],
	}, nil
}

func fetch[T any](ctx context.Context, url string) (*T, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request %s: %w", url, err)
	}
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("unable to fetch %s: %w", url, err)
	}
	defer resp.Body.Close()
	body := new(T)
	if err := json.NewDecoder(resp.Body).Decode(body); err != nil {
		return nil, fmt.Errorf("unable to decode body: %w", err)
	}
	return body, nil
}
