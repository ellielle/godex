package main

import (
	"bufio"
	"fmt"
	"os"
)

const pokeAPIURL = "https://pokeapi.co/api/v2/location-area/"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	prefix := "godex > "

	cliDirections := PokeMap{
		Next:     nil,
		Previous: nil,
		Base:     pokeAPIURL,
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
