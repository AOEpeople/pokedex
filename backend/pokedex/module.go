package pokedex

import (
	"flamingo.me/dingo"
	"flamingo.me/graphql"
)

type Module struct{}

func (*Module) Configure(injector *dingo.Injector) {
	injector.BindMulti(new(graphql.Service)).To(graphqlService{})
}

func (*Module) Depends() []dingo.Module {
	return []dingo.Module{
		new(graphql.Module),
	}
}
