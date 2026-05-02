package main 

import (
	"fmt"
	"errors"
	"math/rand"
)

func commandCatch (cfg * configURL, args []string) error {
	// give user feedback to his input
	if len(args) == 0 {
		return errors.New("Cannot use catch without pokemon argument")
	}
	fmt.Printf("Throwing a pokeball at %s...\n", args[0])
	// get the Stats for the called argument 
	pokeStat, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err 
	}
	// get Baseexperience and simulate catch
	catchLuck := rand.Intn(450)
	if pokeStat.BaseExperience >= catchLuck {
		fmt.Printf("%s escaped!\n", pokeStat.Name)
	} else { 
		fmt.Printf("%s was caught!\n", pokeStat.Name)
		// add caught Pokemon to Pokedex 
		cfg.pokedex[args[0]] = pokeStat
	}
	return nil
}
