package main

//go:generate rm -f graphql/generated.go schema.graphql
//go:generate go run -tags graphql main.go graphql

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"github.com/aoepeople/pokedex/backend/graphql"
	"github.com/aoepeople/pokedex/backend/pokedex"
)

func main() {
	flamingo.App([]dingo.Module{
		new(graphql.Module),
		new(pokedex.Module),
	})
}
