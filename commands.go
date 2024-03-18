package main

import (
	"fmt"
	"os"

	pokeapi "github.com/ellielle/godex/internal/pokeapi"
)

const PokeAPIURL = "https://pokeapi.co/api/v2/location-area/"

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type PokeMap struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

// Displays all command information when 'help' is entered
func (cfg *PokeMap) commandHelp() error {
	commands := getCliCommands(cfg)
	fmt.Println("Usage:")
	for _, info := range commands {
		fmt.Printf("\n%v:  %v", info.name, info.description)
	}
	fmt.Print("\n\n")

	return nil
}

// Exits the program when 'exit' is entered
func (cfg *PokeMap) commandExit() error {
	os.Exit(0)
	return nil
}

// Retrieves the first or next 20 regions from the PokemonAPI
func (cfg *PokeMap) commandMap() error {
	var err error
	var pokeMap pokeapi.PokeResponse

	if cfg.Next != nil {
		pokeMap, err = pokeapi.PokeMapForwards(cfg.Next)
	} else {
		pokeMap, err = pokeapi.PokeMapForwards(PokeAPIURL)
	}
	if err != nil {
		return err
	}

	cfg.Next = pokeMap.Next
	cfg.Previous = pokeMap.Previous

	for _, result := range pokeMap.Results {
		fmt.Println(result.Name)
	}

	return nil
}

// Retrieves the previous 20 regions from the PokemonAPI
func (cfg *PokeMap) commandMapBack() error {
	var err error
	var pokeMap pokeapi.PokeResponse

	if cfg.Previous != nil {
		pokeMap, err = pokeapi.PokeMapBackwards(cfg.Previous)
	} else {
		pokeMap, err = pokeapi.PokeMapBackwards(PokeAPIURL)
	}
	if err != nil {
		return err
	}

	cfg.Next = pokeMap.Next
	cfg.Previous = pokeMap.Previous

	for _, result := range pokeMap.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func getCliCommands(cfg *pokeapi.PokeMap) map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "display a help list for commands",
			callback:    cfg.commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit the Pokedex",
			callback:    cfg.commandExit,
		},
		"map": {
			name:        "map",
			description: "shows 20 regions from the Pokemon world",
			callback:    cfg.commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "shows the previous 20 regions",
			callback:    cfg.commandMapBack,
		},
	}
}
