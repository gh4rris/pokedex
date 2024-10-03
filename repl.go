package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gh4rris/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		var area string
		if commandName == "explore" && len(words) >= 2 {
			area = words[1]
		}
		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(cfg, area)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(words string) []string {
	lower := strings.ToLower(words)
	split := strings.Fields(lower)
	return split
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Go back to the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "lists the pokemon in a given area",
			callback:    commandExplore,
		},
	}
}
