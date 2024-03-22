package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const locationAPIURL = "https://pokeapi.co/api/v2/location-area/"
const pokemonAPIURL = "https://pokeapi.co/api/v2/pokemon/"

// Start a REPL for the program to take user input
func startREPL(cfg *MapConfig) {
	reader := bufio.NewScanner(os.Stdin)
	prefix := "godex > "

	// Continue prompting until the user enters "exit"
	for reader.Text() != "exit" {
		fmt.Print(prefix)
		reader.Scan()
		input := sanitizeInput(reader.Text())

		// Base command will be the first entered word
		commandName := input[0]
		command, ok := getCliCommands()[commandName]
		if !ok {
			fmt.Printf("%v is not a valid command\n", reader.Text())
			continue
		}
		command.callback(cfg, input)
	}
	// If the reader fails, print out the error
	if err := reader.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading stdin: ", err)
	}
}

// Sanitizes input to remove leading and trailing whitespace and convert to lower case to match against getCliCommands()
func sanitizeInput(input string) []string {
	output := strings.Trim(input, " ")
	words := strings.Fields(strings.ToLower(output))
	return words
}
