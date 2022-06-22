package pokedex

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"flamingo.me/graphql"
	"github.com/stretchr/testify/assert"
)

func TestGraphql(t *testing.T) {
	service := new(graphqlService)
	schema, err := os.ReadFile("schema.graphql")
	assert.NoError(t, err)
	assert.Equal(t, schema, service.Schema())

	assert.NotPanics(t, func() { service.Types(new(graphql.Types)) })
}

func TestResolver(t *testing.T) {
	mux := http.NewServeMux()
	testserver := httptest.NewServer(mux)
	defer testserver.Close()

	mux.HandleFunc("/pokemon", func(w http.ResponseWriter, r *http.Request) {
		amount := 151
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		if limit > 0 {
			amount = limit
		}
		pokemonList := make([]string, amount)
		for i := range pokemonList {
			pokemonList[i] = fmt.Sprintf(`{"name": "test%d", "url": "%s/pokemon/test?id=%d"}`, offset+i, testserver.URL, offset+i)
		}
		fmt.Fprintf(w, `{"results": [%s]}`, strings.Join(pokemonList, ","))
	})
	mux.HandleFunc("/pokemon/test", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		fmt.Fprintf(w, `{"id": %s, "name": "test%s", "types": [{"type": {"name": "testtype"}}]}`, id, id)
	})

	resolver := new(Resolver)
	resolver.Inject(&struct {
		BaseURL string "inject:\"config:pokedex.baseurl\""
	}{BaseURL: testserver.URL})

	t.Run("Pokemon", func(t *testing.T) {
		resolver.baseURL = "-://"
		_, err := resolver.Pokemon(context.Background(), nil, nil)
		assert.Error(t, err)
		resolver.baseURL = testserver.URL

		pokemon, err := resolver.Pokemon(context.Background(), nil, nil)
		assert.NoError(t, err)
		assert.Len(t, pokemon, 151)

		pokemon, err = resolver.Pokemon(context.Background(), []int{7, 3, 5}, nil)
		assert.NoError(t, err)
		assert.Len(t, pokemon, 3)
		assert.Equal(t, 3, pokemon[0].ID)
		assert.Equal(t, 5, pokemon[1].ID)
		assert.Equal(t, 7, pokemon[2].ID)

		var catched = true
		pokemon, err = resolver.Pokemon(context.Background(), nil, &catched)
		assert.NoError(t, err)
		assert.Len(t, pokemon, 0)

		catched = false
		pokemon, err = resolver.Pokemon(context.Background(), nil, &catched)
		assert.NoError(t, err)
		assert.Len(t, pokemon, 151)
	})

	t.Run("Total", func(t *testing.T) {
		total, err := resolver.Total(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, total, 151)
	})

	t.Run("TotalCatched", func(t *testing.T) {
		total, err := resolver.TotalCatched(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, total, 0)
	})

	t.Run("SetCatched", func(t *testing.T) {
		_, err := resolver.SetCatched(context.Background(), 0, true)
		assert.NoError(t, err)
	})
}
