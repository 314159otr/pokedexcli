package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
type cliCommand struct {
	name 		string
	description string
	callback 	func() error
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cliCommands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanInput := cleanInput(input)
		if len(cleanInput) == 0 {
			continue
		}
		word := cleanInput[0]
		cliCommand, exists := cliCommands[word]
		if exists {
			cliCommand.callback()
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name: 		 "exit",
			description: "Exit the Pokedex",
			callback: 	 commandExit,
		},
		"help": {
			name: 		 "help",
			description: "Displays a help message",
			callback: 	 commandHelp,
		},
	}
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s", command.name, command.description)
		fmt.Println()
	}
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)

	return words
}
