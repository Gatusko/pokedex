package internal

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Client struct {
	cache Cache
	Pokedex
}

func (c *Client) NewClient(cache Cache) {
	c.cache = cache
	c.Pokedex = NewPokedex()
}

func (c *Client) GetAreas(currentUrl string) (Areas, error) {
	bodyCache, isOnCache := c.cache.get(currentUrl)
	if isOnCache {
		areas := Areas{}
		errM := json.Unmarshal(bodyCache, &areas)
		if errM != nil {
			return Areas{}, errM
		}
		return areas, nil
	}
	res, err := http.Get(currentUrl)
	if err != nil {
		return Areas{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Areas{}, err
	}
	c.cache.add(currentUrl, body)
	areas := Areas{}
	errM := json.Unmarshal(body, &areas)
	if errM != nil {
		return Areas{}, errM
	}
	return areas, nil
}

func (c *Client) ExplorePokemon(url string) (ExplorePokemon, error) {
	bodyCache, isOnCache := c.cache.get(url)
	explorePokemon := ExplorePokemon{}
	if isOnCache {
		errM := json.Unmarshal(bodyCache, &explorePokemon)
		if errM != nil {
			return ExplorePokemon{}, errM
		}
		return explorePokemon, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return ExplorePokemon{}, err
	}
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return explorePokemon, err
	}
	err = json.Unmarshal(body, &explorePokemon)
	if err != nil {
		return explorePokemon, err
	}
	c.cache.add(url, body)
	return explorePokemon, nil
}

func (c *Client) GetPokemon(url string) (PokemonData, error) {
	bodyCache, isOnCache := c.cache.get(url)
	pokemonData := PokemonData{}
	if isOnCache {
		errM := json.Unmarshal(bodyCache, &pokemonData)
		if errM != nil {
			return pokemonData, errM
		}
		return pokemonData, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return pokemonData, err
	}
	if res.StatusCode != 200 {
		return pokemonData, errors.New("Pokemon not found")
	}
	body, err := io.ReadAll(res.Body)
	json.Unmarshal(body, &pokemonData)
	return pokemonData, nil
}
