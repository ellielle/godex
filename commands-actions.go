package main

import (
	"fmt"
	"math/rand"

	"github.com/ellielle/godex/internal/pokedex"
)

// Explore a specific region entered as a parameter to the "explore" command
// Returns Pokemon able to be caught in that region
func commandExplore(cfg *MapConfig, cmd []string) error {
	if len(cmd) != 2 {
		fmt.Println("You must enter a region to explore!")
		return nil
	}

	// Append parameter to the URL to explore that region, throw an error if it fails
	apiURL := locationAPIURL + cmd[1]
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

func commandCatch(cfg *MapConfig, cmd []string) error {
	if len(cmd) != 2 {
		fmt.Println("You must enter a Pokemon to attempt to catch!")
		return nil
	}

	// Append parameter to the URL to attempt to catch that Pokemon, print an error if the command is erroneous
	apiURL := pokemonAPIURL + cmd[1]
	pokemon, err := cfg.Client.PokemonData(apiURL)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Let user know the capture attempt is happening
	fmt.Println("")
	fmt.Printf("Throwing a pokeball at %s...", pokemon.Name)
	fmt.Println("")

	// Use baseEXP as a pseudo capture rate. A successful capture happens when the random number comes back
	// between the "capture range". Higher base exp makes for a harder to catch Pokemon
	baseEXP := pokemon.BaseExperience
	captureRoll := rand.Intn(baseEXP)

	// NOTE: temporary capture formula
	if captureRoll < 50 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("")

		// Save Pokemon data to user's Pokedex
		cfg.Pokedex.AddPokemon(pokemon)
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemon.Name)
	fmt.Println("")
	return nil
}

func commandInspect(cfg *MapConfig, cmd []string) error {
	if len(cmd) != 2 {
		fmt.Println("You must enter a Pokemon to inspect!")
		return nil
	}

	// Get pokemon's name from user input
	pokemon := cmd[1]

	// Attempt to get information from the Pokedex if the Pokemon has been caught before
	// Otherwise, inform the user they haven't caught that Pokemon yet
	dat, ok := cfg.Pokedex.Entries[pokemon]
	if !ok {
		fmt.Println("You have not caught that Pokemon yet!")
		return nil
	}

	printStats(&dat)
	return nil
}

// Prints the Pokemon's stats out
func printStats(pokemon *pokedex.Pokemon) {
	fmt.Printf("\nInspecting %s...", pokemon.Name)
	fmt.Printf("\nHeight: %v", pokemon.Height)
	fmt.Printf("\nWeight: %v", pokemon.Weight)
	fmt.Println("")
	fmt.Println("\nStats:")
	for _, stats := range pokemon.Stats {
		fmt.Printf("%s: %v\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Println("\nTypes:")
	for _, types := range pokemon.Types {
		fmt.Printf("%s\n", types.Type.Name)
	}
	fmt.Println("")
}
