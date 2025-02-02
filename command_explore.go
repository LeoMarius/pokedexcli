package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {

	if len(args) < 1 {
		fmt.Println("Localisation manquante")
		return fmt.Errorf("usage: explore <location>")
	}

	locationsResp, err := cfg.pokeapiClient.ExploreLocations(cfg.nextLocationsURL, args[0])
	if err != nil {
		fmt.Println("Espace inconnu")
		return err
	}

	if len(locationsResp.PokemonEncounters) == 0 {
		fmt.Println("Aucun pokémon trouvé")
		return nil
	}

	fmt.Printf("%v Pokémons trouvés à %v", len(locationsResp.PokemonEncounters), args[0])
	fmt.Println("")

	for _, loc := range locationsResp.PokemonEncounters {
		fmt.Println(loc.Pokemon.Name)
	}

	fmt.Println(args[0])

	// fmt.Println(locationsResp)

	// cfg.nextLocationsURL = locationsResp.Next
	// cfg.prevLocationsURL = locationsResp.Previous

	// for _, loc := range locationsResp.Results {
	// 	fmt.Println(loc.Name)
	// }
	return nil
}
