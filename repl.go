package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const pokeAPIURL = "https://pokeapi.co/api/v2/location-area/"

func startREPL(cfg *PokeMap, cliCommands map[string]cliCommand) {
	reader := bufio.NewScanner(os.Stdin)
	prefix := "godex > "

	// Continue prompting until the user enters "exit"
	for reader.Text() != "exit" {
		fmt.Print(prefix)
		reader.Scan()
		input := sanitizeInput(reader.Text())
		commandName := input[0]
		// Check for the entered command in the command list
		command, ok := cliCommands[commandName]
		if !ok {
			fmt.Printf("%v is not a valid command\n", reader.Text())
			continue
		}
		command.callback(cfg)
	}
	if err := reader.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading stdin: ", err)
	}
}

func sanitizeInput(input string) []string {
	output := strings.Trim(input, " ")
	words := strings.Fields(strings.ToLower(output))
	return words
}
