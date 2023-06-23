package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

type Areas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
type Client struct {
	cache Cache
}

func (c *Client) NewClient(cache Cache) {
	c.cache = cache
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
	c.cache.add(currentUrl, body)
	if err != nil {
		return Areas{}, err
	}
	areas := Areas{}
	errM := json.Unmarshal(body, &areas)
	if errM != nil {
		return Areas{}, errM
	}
	return areas, nil
}
