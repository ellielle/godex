package main

import (
	"fmt"
)

func main() {
	cliDirections := PokeMap{
		Next:     nil,
		Previous: nil,
		Base:     pokeAPIURL,
	}
	cliCommands := getCliCommands(&cliDirections)
	fmt.Println("Welcome to Godex, a CLI Pokedex in Go!")
	startREPL(cliCommands)
}
