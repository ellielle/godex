package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*MapConfig, []string) error
}

// Displays all command information when 'help' is entered
func commandHelp(cfg *MapConfig, cmd []string) error {
	fmt.Println("Usage:")
	for _, command := range getCliCommands() {
		fmt.Printf("\n%v:  %v", command.name, command.description)
	}
	fmt.Print("\n\n")

	return nil
}

// Exits the program when 'exit' is entered
func commandExit(cfg *MapConfig, cmd []string) error {
	os.Exit(0)
	return nil
}

// Return CLI commands as a map so a command can be accessed by name
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
		"explore": {
			name:        "explore map",
			description: "explores user-specified map",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch pokemon",
			description: "catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect pokemon",
			description: "inspect a pokemon's height, weight, stats and type(s)",
			callback:    commandInspect,
		},
	}
}
