package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("Your Pokedex is empty. Go catch some Pokemon!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range cfg.pokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}