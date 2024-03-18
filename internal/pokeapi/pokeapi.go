package pokeapi

import (
	"encoding/json"
	"net/http"
)

const PokeAPIURL = "https://pokeapi.co/api/v2/location-area/"

type PokeResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// Retrieves 20 results from the location area section of the PokeAPI
func PokeMapMove(apiUrl string) (PokeResponse, error) {
	var res *http.Response
	var err error

	res, err = http.Get(apiUrl)

	decoder := json.NewDecoder(res.Body)
	pokeMap := PokeResponse{}
	err = decoder.Decode(&pokeMap)
	if err != nil {
		return PokeResponse{}, err
	}

	return pokeMap, nil
}
