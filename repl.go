package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"github.com/314159otr/pokedexcli/internal/pokeapi"
)
type cliCommand struct {
	name 		string
	description string
	callback 	func(*config) error
}

type config struct {
	client	 pokeapi.Client
	next	 *string
	previous *string
}

func startRepl(config *config) {
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
			cliCommand.callback(config)
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
		"map": {
			name: 		 "map",
			description: "Displays the names of the next 20 location ares in the Pokemon World.",
			callback: 	 commandMap,
		},
		"mapb": {
			name: 		 "mapb",
			description: "Displays the names of the previous 20 location ares in the Pokemon World.",
			callback: 	 commandMapb,
		},
	}
}

func commandMapb(config *config) error {
	if config.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locationAreasResponse, err := config.client.GetLocationAreas(config.previous)
	if err != nil {
		return err
	}
	config.next = locationAreasResponse.Next
	config.previous = locationAreasResponse.Previous
	for _, locationArea := range locationAreasResponse.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}
func commandMap(config *config) error {
	if config.previous != nil && config.next == nil {
		fmt.Println("you're on the last page")
		return nil
	}
	locationAreasResponse, err := config.client.GetLocationAreas(config.next)
	if err != nil {
		return err
	}
	config.next = locationAreasResponse.Next
	config.previous = locationAreasResponse.Previous
	for _, locationArea := range locationAreasResponse.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s", command.name, command.description)
		fmt.Println()
	}
	return nil
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)

	return words
}
