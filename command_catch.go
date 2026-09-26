package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemon := args[0]
	resp, err := cfg.pokeapiClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	caught := catchPokemon(resp.BaseExperience)
	if caught {
		fmt.Printf("%s was caught!\n", pokemon)
		cfg.caughtPokemon[pokemon] = resp
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}

func catchPokemon(baseExperience int) bool {
	chance := max(100-baseExperience/3, 5)

	roll := rand.Intn(100) + 1

	return roll <= chance
}
