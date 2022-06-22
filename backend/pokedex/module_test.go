package pokedex

import (
	"testing"

	"flamingo.me/flamingo/v3/framework/config"
	"github.com/stretchr/testify/assert"
)

func TestModule(t *testing.T) {
	assert.NoError(t, config.TryModules(config.Map{"pokedex.baseurl": "http://example.com"}, new(Module)))
}
