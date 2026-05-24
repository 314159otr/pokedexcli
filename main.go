package main

import (
	"time"
	"github.com/314159otr/pokedexcli/internal/pokeapi"
)
func main() {
	timeout := 5 * time.Second
	client := pokeapi.NewClient(timeout)
	config := &config{
		client: client,
	}
	startRepl(config)
}

