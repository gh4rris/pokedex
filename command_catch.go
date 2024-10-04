package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("pokemon name required")
	}
	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	rng := rand.Intn(pokemon.BaseExperience)
	if rng <= 50 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		if _, ok := cfg.caughtPokemon[pokemon.Name]; !ok {
			cfg.caughtPokemon[pokemon.Name] = pokemon
			fmt.Printf("%s's data was added to the pokedex\n", pokemon.Name)
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}
