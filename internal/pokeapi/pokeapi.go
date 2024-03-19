package pokeapi

import (
	"encoding/json"
	"fmt"
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
func (c *Client) PokeMapNext(apiUrl string, next *string) (PokeResponse, error) {
	var res *http.Response
	var err error

	if next != nil {
		res, err = http.Get(*next)
	} else {
		res, err = http.Get(apiUrl)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	pokeMap := PokeResponse{}
	err = decoder.Decode(&pokeMap)
	if err != nil {
		return PokeResponse{}, err
	}

	return pokeMap, nil
}

func PokeMapPrevious(previous *string) (PokeResponse, error) {
	var res *http.Response
	var err error

	if previous != nil {
		res, err = http.Get(*previous)
	} else {
		fmt.Println("There are no previous regions to display!")
		return PokeResponse{}, nil
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	pokeMap := PokeResponse{}
	err = decoder.Decode(&pokeMap)
	if err != nil {
		return PokeResponse{}, err
	}

	return pokeMap, nil
}
