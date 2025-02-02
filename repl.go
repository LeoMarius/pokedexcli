package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/LeoMarius/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)

			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", commandName)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the map",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the 20 previous map",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore la location - 1 argument",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Capture un pokemon - 1 argument",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect un pokemon - 1 argument",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Donne la liste des pokemons capturés",
			callback:    openPokedex,
		},
	}
	return commands
}
