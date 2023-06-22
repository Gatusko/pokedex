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

func GetAreas(currentUrl string) (Areas, error) {
	res, err := http.Get(currentUrl)
	if err != nil {
		return Areas{}, err
	}
	body, err := io.ReadAll(res.Body)
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
