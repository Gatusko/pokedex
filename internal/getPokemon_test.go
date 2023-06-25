package internal

import (
	"testing"
	"time"
)

func TestGetPokemon(t *testing.T) {
	cache := Cache{}
	cache.NewCache(10 * time.Second)
	client := Client{}
	client.NewClient(cache)
	_, err := client.ExplorePokemon("https://pokeapi.co/api/v2/location-area/canalave-city-area")
	if err != nil {
		return
	}
	t.Errorf("Expected to find the area")
	_, err = client.ExplorePokemon("https://pokeapi.co/api/v2/location-area/canalave-city-areaewfaerfa")
	if err != nil {
		t.Errorf("Expected to not find any pokemon")
	}
	return
}
