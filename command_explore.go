package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location area name")
	}
	name := args[0]

	fmt.Printf("Exploring %s...\n", name)
	area, err := fetchLocationArea(name)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range area.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}