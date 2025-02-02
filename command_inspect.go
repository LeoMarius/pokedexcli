package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {

	if len(args) < 1 {
		fmt.Println("Nobody to inspect")
		return fmt.Errorf("usage: inspect <pokemon>")
	}

	p, err := cfg.pokeapiClient.InspectPokemon(args[0])
	if err != nil {
		fmt.Println("you have not caught that pokemon")

		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v\n", p.Name)
	fmt.Printf("Weight: %v\n", p.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range p.Stats {

		fmt.Printf(" - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, types := range p.Types {
		fmt.Printf(" - %v\n", types.Type.Name)
	}

	return nil
}

func openPokedex(cfg *config, args ...string) error {

	pokedex, err := cfg.pokeapiClient.GetPokedex()
	if err != nil {
		return fmt.Errorf("Il y a un problem avec le pokedex")
	}

	if len(pokedex) == 0 {
		fmt.Println("Your Pokedex is empty")
		return nil
	}

	fmt.Println("Your Pokedex :")

	for _, pokemon := range pokedex {
		fmt.Printf(" - %v\n", pokemon.Name)
	}

	return nil

}
