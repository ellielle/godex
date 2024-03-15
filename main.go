package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type Config struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

// TODO: PokeResponse will be used in the API fetch JSON marshaling
// TODO: may or may not need the separate Config struct?
type PokeResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const PokeAPIURL = "https://pokeapi.co/api/v2/location-area/"

// Displays all command information when 'help' is entered
func (cfg *Config) commandHelp() error {
	commands := getCliCommands(cfg)
	fmt.Println("Usage:")
	for _, info := range commands {
		fmt.Printf("\n%v:  %v", info.name, info.description)
	}
	fmt.Print("\n\n")

	return nil
}

// Exits the program when 'exit' is entered
func (cfg *Config) commandExit() error {
	os.Exit(0)
	return nil
}

// Retrieves the first or next 20 regions from the PokemonAPI
func (cfg *Config) commandMap() error {
	//  response := Config{}
	res, err := http.Get(PokeAPIURL)
	fmt.Printf("RESPONSE: %v", res)
	if err != nil {
		return err
	}

	return nil
}

// Retrieves the previous 20 regions from the PokemonAPI
func (cfg *Config) commandMapBack() error {
	return nil
}

func getCliCommands(cfg *Config) map[string]cliCommand {
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
	cliConfig := Config{
		Next:     nil,
		Previous: nil,
	}
	cliCommands := getCliCommands(&cliConfig)

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
