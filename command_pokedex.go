package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	if len(cfg.caughtPokemon) == 0 {
		return fmt.Errorf("you haven't caught any pokemon")
	}
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}
