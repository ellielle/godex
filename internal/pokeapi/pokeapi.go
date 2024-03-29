package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ellielle/godex/internal/pokedex"
)

type LocationResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokeResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// Retrieves 20 results from either the Next, or if available, Previous locations URL
func (c *Client) ListMapLocations(apiURL string) (LocationResponse, error) {
	locationResp := LocationResponse{}

	// Check for an entry in the cache before requesting
	val, ok := c.cache.Get(apiURL)
	if ok {
		err := json.Unmarshal([]byte(val), &locationResp)
		if err != nil {
			return LocationResponse{}, err
		}
		return locationResp, nil
	}

	// Create a request to the API and then send it
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return LocationResponse{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResponse{}, err
	}
	defer resp.Body.Close()

	// Read the response body, and then Unmarshal it into  LocationResponse
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationResponse{}, err
	}

	// Unmarshal JSON data into response struct
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return LocationResponse{}, err
	}

	// Add URL and data to the cache
	c.cache.Add(apiURL, &dat)
	return locationResp, nil
}

func (c *Client) ListPokemon(apiURL string) (PokeResponse, error) {
	pokeResp := PokeResponse{}

	// Check for an entry in the cache before requesting
	val, ok := c.cache.Get(apiURL)
	if ok {
		err := json.Unmarshal([]byte(val), &pokeResp)
		if err != nil {
			return PokeResponse{}, err
		}
		return pokeResp, nil
	}

	// Create a request to the API and then send it
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return PokeResponse{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeResponse{}, err
	}
	defer resp.Body.Close()

	// Read the body from the request
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeResponse{}, err
	}

	// Unmarshal JSON data into response struct
	err = json.Unmarshal(dat, &pokeResp)
	if err != nil {
		return PokeResponse{}, err
	}

	// Add URL and data to the cache
	c.cache.Add(apiURL, &dat)
	return pokeResp, nil
}

func (c *Client) PokemonData(apiURL string) (pokedex.Pokemon, error) {
	pokeResp := pokedex.Pokemon{}

	// Check for an entry in the cache before requesting
	val, ok := c.cache.Get(apiURL)
	if ok {
		err := json.Unmarshal([]byte(val), &pokeResp)
		if err != nil {
			return pokedex.Pokemon{}, err
		}
		return pokeResp, nil
	}

	// Create a request to the API and then send it
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return pokedex.Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokedex.Pokemon{}, err
	}
	defer resp.Body.Close()

	// Read the body from the request
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokedex.Pokemon{}, err
	}

	// Unmarshal JSON data into response struct
	err = json.Unmarshal(dat, &pokeResp)
	if err != nil {
		return pokedex.Pokemon{}, err
	}

	// Add URL and data to the cache
	c.cache.Add(apiURL, &dat)
	return pokeResp, nil
}
