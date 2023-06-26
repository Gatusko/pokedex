package internal

import (
	"testing"
	"time"
)

func TestPokemonData(t *testing.T) {
	cache := Cache{}
	cache.NewCache(10 * time.Second)
	client := Client{}
	client.NewClient(cache)
	_, err := client.GetPokemon("https://pokeapi.co/api/v2/pokemon/mewtwo")
	if err != nil {
		t.Errorf("Expectec to get a pokemon")
	}

	_, err = client.GetPokemon("https://pokeapi.co/api/v2/pokemon/qwert")
	if err != nil {
		return
	}
	t.Errorf("Expectec to not get a Pokemon")
	return
}
