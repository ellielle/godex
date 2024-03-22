package main

import (
	"fmt"
	"time"

	"github.com/ellielle/godex/internal/pokeapi"
)

func main() {
	cliDirections := &PokeMap{
		Next:     nil,
		Previous: nil,
		Client:   pokeapi.NewClient(5*time.Second, (60*5)*time.Second),
	}
	cliCommands := getCliCommands()
	fmt.Println("Welcome to Godex, a CLI Pokedex in Go!")
	startREPL(cliDirections, cliCommands)
}
