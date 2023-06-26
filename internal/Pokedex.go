package internal

import (
	"fmt"
	"math/rand"
)

type Pokedex struct {
	pokemonCatched map[string]PokemonData
}

func (p *Pokedex) listAllPokemon() {
	for _, pokemon := range p.pokemonCatched {
		fmt.Println(pokemon)
	}
}

func (p *Pokedex) addPokemon(pokemon PokemonData) {
	p.pokemonCatched[pokemon.Name] = pokemon
}

func (p *Pokedex) GetPokemon(namePokemon string) (PokemonData, bool) {
	pokemonData, ok := p.pokemonCatched[namePokemon]
	if !ok {
		return pokemonData, false
	}
	return pokemonData, true
}

func (p *Pokedex) CatchPokemon(pokemon PokemonData) bool {
	intExperience := int(pokemon.BaseExperience)
	if isCached(intExperience) {
		p.addPokemon(pokemon)
		return true
	} else {
		return false
	}

}

func isCached(baseExperience int) bool {
	randomInt := rand.Intn(baseExperience)
	return randomInt >= baseExperience/2
}

func NewPokedex() Pokedex {
	return Pokedex{make(map[string]PokemonData)}
}
