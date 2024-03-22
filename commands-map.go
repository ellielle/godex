package main

import (
	"fmt"

	"github.com/ellielle/godex/internal/pokeapi"
)

type MapConfig struct {
	Next     *string
	Previous *string
	Client   pokeapi.Client
}

// Retrieves the first or next 20 regions from the PokemonAPI
func commandMap(cfg *MapConfig, cmd []string) error {
	apiURL := pokeAPIURL
	if cfg.Next != nil {
		apiURL = *cfg.Next
	}
	pokeMap, err := cfg.Client.ListMapLocations(apiURL)
	if err != nil {
		return nil
	}

	cfg.Next = pokeMap.Next
	cfg.Previous = pokeMap.Previous

	fmt.Println("")
	for _, result := range pokeMap.Results {
		fmt.Println(result.Name)
	}
	fmt.Println("")

	return nil
}

// Retrieves the previous 20 regions from the PokemonAPI
func commandMapBack(cfg *MapConfig, cmd []string) error {
	// Print an error letting the user know there are no previous locations yet
	if cfg.Previous == nil {
		fmt.Println("")
		fmt.Println("There are no previous regions to display!")
		fmt.Println("")
		return nil
	}

	apiURL := *cfg.Previous
	pokeMap, err := cfg.Client.ListMapLocations(apiURL)
	if err != nil {
		return err
	}

	cfg.Next = pokeMap.Next
	cfg.Previous = pokeMap.Previous

	fmt.Println("")
	for _, result := range pokeMap.Results {
		fmt.Println(result.Name)
	}
	fmt.Println("")

	return nil
}

// Explore a specific region entered as a parameter to the "explore" command
// Returns Pokemon able to be caught in that region
func commandExplore(cfg *MapConfig, cmd []string) error {
	if len(cmd) != 2 {
		fmt.Println("You must enter a region to explore!")
		return nil
	}

	// Append parameter to the URL to explore that region, throw an error if it fails
	apiURL := pokeAPIURL + cmd[1]
	pokemonList, err := cfg.Client.ListPokemon(apiURL)
	if err != nil {
		fmt.Println("Invalid region! Please try again.")
		return err
	}

	// Print all pokemon in the region to the CLI
	fmt.Println("")
	for _, pk := range pokemonList.PokemonEncounters {
		fmt.Println(pk.Pokemon.Name)
	}
	fmt.Println("")
	return nil
}
