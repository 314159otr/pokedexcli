package main

import (
	"time"
	"github.com/314159otr/pokedexcli/internal/pokeapi"
)
func main() {
	timeout := 5 * time.Second
	cacheInterval :=  5 * time.Minute
	client := pokeapi.NewClient(timeout, cacheInterval)
	config := &config{
		client: client,
		pokedex: map[string]pokeapi.PokemonResponse{},
	}
	startRepl(config)
}

