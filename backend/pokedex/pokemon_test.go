package pokedex

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokemon(t *testing.T) {
	mux := http.NewServeMux()
	testserver := httptest.NewServer(mux)
	defer testserver.Close()

	mux.HandleFunc("/pokemonlist", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"results": [{"name": "test1", "url": "%s/pokemon1"}, {"name": "test2", "url": "%s/pokemon2"}]}`, testserver.URL, testserver.URL)
	})
	mux.HandleFunc("/brokenpokemonlist", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"results": [{"name": "test", "url": "-://"}]}`)
	})
	mux.HandleFunc("/pokemon1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"id": 1, "name": "test1", "types": [{"type": {"name": "testtype"}}]}`)
	})
	mux.HandleFunc("/pokemon2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"id": 2, "name": "test2", "types": [{"type": {"name": "testtype"}}]}`)
	})
	mux.HandleFunc("/clienterror", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusContinue)
		conn, _, _ := w.(http.Hijacker).Hijack()
		conn.Close()
	})

	t.Run("fetch", func(t *testing.T) {
		type testResult struct{}

		_, err := fetch[testResult](context.Background(), "-://")
		assert.Error(t, err)

		_, err = fetch[testResult](context.Background(), testserver.URL+"/clienterror")
		assert.Error(t, err)

		_, err = fetch[testResult](context.Background(), testserver.URL)
		assert.Error(t, err)

		_, err = fetch[testResult](context.Background(), testserver.URL+"/pokemon1")
		assert.NoError(t, err)
	})

	t.Run("fetchPokemon", func(t *testing.T) {
		_, err := fetchPokemon(context.Background(), "-://")
		assert.Error(t, err)

		pokemon, err := fetchPokemon(context.Background(), testserver.URL+"/pokemon1")
		assert.NoError(t, err)
		assert.Equal(t, "test1", pokemon.Name)
		assert.False(t, pokemon.Catched)

		setCatched(1)
		pokemon, err = fetchPokemon(context.Background(), testserver.URL+"/pokemon1")
		assert.NoError(t, err)
		assert.Equal(t, "test1", pokemon.Name)
		assert.True(t, pokemon.Catched)

		unsetCatched(1)
		pokemon, err = fetchPokemon(context.Background(), testserver.URL+"/pokemon1")
		assert.NoError(t, err)
		assert.Equal(t, "test1", pokemon.Name)
		assert.False(t, pokemon.Catched)
	})

	t.Run("fetchPokemonList", func(t *testing.T) {
		_, err := fetchPokemonList(context.Background(), "-://")
		assert.Error(t, err)

		_, err = fetchPokemonList(context.Background(), testserver.URL+"/brokenpokemonlist")
		assert.Error(t, err)

		pokemon, err := fetchPokemonList(context.Background(), testserver.URL+"/pokemonlist")
		assert.NoError(t, err)
		assert.Len(t, pokemon, 2)
		assert.Equal(t, "test1", pokemon[0].Name)
		assert.Equal(t, "test2", pokemon[1].Name)
		assert.Equal(t, []string{"testtype"}, pokemon[0].Type)
	})
}
