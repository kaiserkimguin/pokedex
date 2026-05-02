package main 

import (
	"fmt"
)

func commandPokedex (cfg *configURL, args []string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("you have not caught any pokemon so far")
	} else {
		fmt.Println("your pokedex:")
		for pokemon, _  := range cfg.pokedex {
			fmt.Printf("    -%s\n", pokemon)
		}
	}
	return nil

}
