package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("pokemon name required")
	}
	name := args[0]
	if pokemon, ok := cfg.caughtPokemon[name]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokeType := range pokemon.Types {
			fmt.Printf("- %s\n", pokeType.Type.Name)
		}
		return nil
	}
	return fmt.Errorf("you have not caught that pokemon")
}
