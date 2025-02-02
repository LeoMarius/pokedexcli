package main

import (
	"time"

	"github.com/LeoMarius/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	// pokecache := pokecache.NewCache(1)

	startRepl(cfg)
}
