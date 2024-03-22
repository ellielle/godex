package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const pokeAPIURL = "https://pokeapi.co/api/v2/location-area/"

func startREPL(cliCommands map[string]cliCommand) {
	scanner := bufio.NewScanner(os.Stdin)
	prefix := "godex > "

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

func sanitizeInput(input string) []string {
	output := strings.Trim(input, " ")
	// FIXME: strings.Fields - look up
	words := strings.Fields(output)
	return words
}
