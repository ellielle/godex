package main

import (
	"fmt"
)

// Prints out a list of all Pokemon currently in the user's Pokedex
func commandPokedex(cfg *MapConfig, cmd []string) error {
	fmt.Println("\nYour Pokedex:")
	for _, pokemon := range cfg.Pokedex.Entries {
		fmt.Println(" - " + pokemon.Name)
	}
	fmt.Println("")
	return nil
}
