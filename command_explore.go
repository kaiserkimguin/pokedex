package main 

import (
	"fmt"
	"errors"
)

func commandExplore (cfg *configURL, args []string) error {
	// check wether correct args were provided
	if len(args) == 0 {
		return errors.New("no arguments provided, cannot get data")
	}
	// Tell user whats happening
	fmt.Printf("Exploring %s...\n", args[0])
	// call GetLocationArea on Client to recieve encounters
	explResp, err := cfg.pokeapiClient.GetLocationArea(args[0])
	if err != nil {
		return err 
	}
	// printing all possible encounters of area 
	fmt.Println("Found Pokemon:")
	for _, enc := range explResp.PokemonEncounters {
		fmt.Println(enc.Pokemon.Name)
	}
	return nil
}


