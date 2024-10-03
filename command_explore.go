package main

import (
	"fmt"
)

func commandExplore(cfg *config, area string) error {
	fmt.Printf("Exploring %v...\n", area)
	resp, err := cfg.pokeapiClient.PokemonTypes(area)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
