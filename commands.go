package main

import (
	"fmt"
	"os"

	"github.com/ellielle/godex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*PokeMap) error
}

type PokeMap struct {
	Next     *string
	Previous *string
	Client   pokeapi.Client
}

// Displays all command information when 'help' is entered
func commandHelp(cfg *PokeMap) error {
	fmt.Println("Usage:")
	for _, command := range getCliCommands() {
		fmt.Printf("\n%v:  %v", command.name, command.description)
	}
	fmt.Print("\n\n")

	return nil
}

// Exits the program when 'exit' is entered
func commandExit(cfg *PokeMap) error {
	os.Exit(0)
	return nil
}

// Retrieves the first or next 20 regions from the PokemonAPI
func commandMap(cfg *PokeMap) error {
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

	for _, result := range pokeMap.Results {
		fmt.Println(result.Name)
	}

	return nil
}

// Retrieves the previous 20 regions from the PokemonAPI
func commandMapBack(cfg *PokeMap) error {
	apiURL := pokeAPIURL
	if cfg.Previous != nil {
		apiURL = *cfg.Previous
	}
	pokeMap, err := cfg.Client.ListMapLocations(apiURL)
	if err != nil {
		fmt.Println("There are no previous regions to display!")
		return err
	}

	cfg.Next = pokeMap.Next
	cfg.Previous = pokeMap.Previous

	for _, result := range pokeMap.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "display a help list for commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "shows 20 regions from the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "shows the previous 20 regions",
			callback:    commandMapBack,
		},
	}
}
