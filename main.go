package main

import (
	"fmt"
	"time"

	"github.com/ellielle/godex/internal/pokeapi"
)

func main() {
	cliDirections := &MapConfig{
		Next:     nil,
		Previous: nil,
		Client:   pokeapi.NewClient(5*time.Second, (60*5)*time.Second),
	}
	fmt.Println("Welcome to Godex, a CLI Pokedex in Go!")
	startREPL(cliDirections)
}
