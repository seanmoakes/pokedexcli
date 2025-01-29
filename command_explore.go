package main

import (
	"fmt"
)

func commandExplore(cfg *config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("no area entered")
	}

	name := params[0]
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", name)
	fmt.Println("Found Pokemon:")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
