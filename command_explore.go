package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Location name required")
	}
	area := args[0]
	location, err := cfg.pokeapiClient.LocationIDInfo(area)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %v...\n", location.Name)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
