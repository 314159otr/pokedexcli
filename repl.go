package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"math/rand"
	"github.com/314159otr/pokedexcli/internal/pokeapi"
)
type cliCommand struct {
	name 		string
	description string
	callback 	func(*config, []string) error
}

type config struct {
	client	 pokeapi.Client
	next	 *string
	previous *string
	pokedex  map[string]pokeapi.PokemonResponse
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
			cliCommand.callback(config, cleanInput[1:])
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
		"explore": {
			name: 		 "explore <location_name>",
			description: "Displays the Pokemon from a location area",
			callback: 	 commandExplore,
		},
		"catch": {
			name: 		 "catch <pokemon_name>",
			description: "Tries to catch a pokemon",
			callback: 	 commandCatch,
		},
		"inspect": {
			name: 		 "inspect <pokemon_name>",
			description: "Displays the pokemon details",
			callback: 	 commandInspect,
		},
		"pokedex": {
			name: 		 "pokedex",
			description: "Displays all pokemon the user has caught",
			callback: 	 commandPokedex,
		},
	}
}

func commandPokedex(config *config, args []string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.pokedex {
		fmt.Printf(" - %v\n", pokemon.Name)
	}
	return nil
}

func commandInspect(config *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("No pokemon specified. Usage: catch <pokemon_name>")
		return fmt.Errorf("No pokemon specified. Usage: catch <pokemon_name>")
	}
	pokemon, found := config.pokedex[args[0]]
	if !found {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  -%v\n", pokemonType.Type.Name)
	}
	return nil
}

func commandCatch(config *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("No pokemon specified. Usage: catch <pokemon_name>")
		return fmt.Errorf("No pokemon specified. Usage: catch <pokemon_name>")
	}

	pokemonResponse, err := config.client.GetPokemonByName(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResponse.Name)
	chance := 100.0 / (100.0 + float64(pokemonResponse.BaseExperience))
	caught := rand.Float64() < chance
	if !caught {
		fmt.Printf("%s escaped!\n", pokemonResponse.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemonResponse.Name)
	config.pokedex[pokemonResponse.Name] = pokemonResponse
	return nil
}

func commandExplore(config *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("No location area specified. Usage: explore <location_name>")
		return fmt.Errorf("No location area specified. Usage: explore <location_name>")
	}

	fmt.Println("Exploring " + args[0] + "...")
	locationAreaResponse, err := config.client.GetLocationAreaByName(args[0])
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range locationAreaResponse.PokemonEncounters {
		fmt.Println(" - " + pokemonEncounter.Pokemon.Name)
	}

	return nil
}

func commandMapb(config *config, args []string) error {
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
func commandMap(config *config, args []string) error {
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

func commandHelp(config *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s", command.name, command.description)
		fmt.Println()
	}
	return nil
}

func commandExit(config *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)

	return words
}
