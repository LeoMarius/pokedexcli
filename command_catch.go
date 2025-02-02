package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) < 1 {
		fmt.Println("Throwing a Pokeball at nobody won't catch anything")
		return fmt.Errorf("usage: catch <pokemon>")
	}

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		fmt.Println("Throwing a Pokeball at nobody won't catch anything")
		return fmt.Errorf("usage: catch <pokemon>")
	}

	fmt.Printf("Throwing a Pokeball at %v...", args[0])
	fmt.Println("")

	pokemon_xp := pokemonResp.BaseExperience

	src := rand.NewSource(time.Now().UnixNano()) // Crée une source de nombres pseudo-aléatoires
	rng := rand.New(src)                         // Initialise un générateur avec cette source

	random := rng.Intn(501) // Génère un entier entre 0 et 500 inclus

	if pokemon_xp > 450 {
		pokemon_xp = 450
	}

	fmt.Printf("Jet à %v pour un Pokemon à %v...", random, pokemon_xp)

	if random > pokemon_xp {
		fmt.Println("Capture réussie")
		cfg.pokeapiClient.AddPokemon(pokemonResp)
	} else {
		fmt.Println("Capture échouée")
	}

	return nil
}
