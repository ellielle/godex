package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokeResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// Retrieves 20 results from either the Next, or if available, Previous locations URL
func (c *Client) ListMapLocations(apiURL string) (PokeResponse, error) {
	var err error

	// Check for an entry in the cache before requesting
	val, ok := c.cache.Get(apiURL)
	if ok {
		fmt.Print("\nUSING CACHE\n")
		locationResp := PokeResponse{}
		err = json.Unmarshal([]byte(val), &locationResp)
		if err != nil {
			return PokeResponse{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return PokeResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeResponse{}, err
	}

	locationResp := PokeResponse{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return PokeResponse{}, err
	}

	return locationResp, nil
}
