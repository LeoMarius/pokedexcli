package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	if len(args) < 1 {
		fmt.Println("Localisation manquante")
		return fmt.Errorf("usage: explore <location>")
	}

	locationsResp, err := cfg.pokeapiClient.GetLocation(args[0])
	if err != nil {
		fmt.Println("Espace inconnu")
		return err
	}

	if len(locationsResp.PokemonEncounters) == 0 {
		fmt.Println("Aucun pokémon trouvé")
		return nil
	}

	fmt.Printf("%v Pokémons trouvés à %v\n", len(locationsResp.PokemonEncounters), args[0])
	fmt.Println("")

	for _, enc := range locationsResp.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
