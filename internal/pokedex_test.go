package internal

import "testing"

func TestAddPokemon(t *testing.T) {
	pokedex := NewPokedex()
	var stats []Stats
	stats = append(stats, Stats{1, 1, Stat{"test"}})
	var types []Types
	pokemon := PokemonData{"test", 1.0, 1, stats, 1, types}
	pokedex.addPokemon(pokemon)
}

func TestListAllPokemon(t *testing.T) {
	pokedex := NewPokedex()
	var stats []Stats
	stats = append(stats, Stats{1, 1, Stat{"test"}})
	var types []Types
	pokemon := PokemonData{"test", 1.0, 1, stats, 1, types}
	pokedex.addPokemon(pokemon)
	pokedex.ListAllPokemon()
}

func TestPokedex_CatchPokemon(t *testing.T) {
	pokedex := NewPokedex()
	var stats []Stats
	stats = append(stats, Stats{5, 1, Stat{"test"}})
	var types []Types
	pokemon := PokemonData{"test", 1.0, 1, stats, 1, types}
	for i := 0; i <= 10; i++ {
		pokedex.CatchPokemon(pokemon)
	}
	pokemon, ok := pokedex.GetPokemon("test")
	if !ok {
		t.Errorf("need to found pokemon")
	}
	return
}
