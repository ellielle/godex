package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

type PokeResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
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
	var res *http.Response
	var err error
	if cfg.Next != nil {
		res, err = http.Get(*cfg.Next)
	} else {
		res, err = http.Get(PokeAPIURL)
	}
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(res.Body)
	pokeMap := PokeResponse{}
	err = decoder.Decode(&pokeMap)
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
	var res *http.Response
	var err error
	if cfg.Previous != nil {
		res, err = http.Get(*cfg.Previous)
	} else {
		fmt.Println("You haven't gone anywhere yet!")
		return nil
	}

	decoder := json.NewDecoder(res.Body)
	pokeMap := PokeResponse{}
	err = decoder.Decode(&pokeMap)
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

func getCliCommands(cfg *PokeMap) map[string]cliCommand {
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	prefix := "godex > "
	cliDirections := PokeMap{
		Next:     nil,
		Previous: nil,
	}
	cliCommands := getCliCommands(&cliDirections)

	fmt.Println("Welcome to Godex, a CLI Pokedex in Go!")

	// Continue prompting until the user enters "exit"
	for scanner.Text() != "exit" {
		fmt.Print(prefix)
		scanner.Scan()
		// Check for the entered command in the command list
		command, ok := cliCommands[scanner.Text()]
		if !ok {
			fmt.Printf("%v is not a valid command\n", scanner.Text())
			continue
		}
		command.callback()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading stdin: ", err)
	}
}
