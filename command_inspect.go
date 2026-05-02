package main 

import (
	"fmt"
)

func commandInspect (cfg *configURL, args []string) error {
	pokedexEntry, exists := cfg.pokedex[args[0]]
	if exists {
		fmt.Printf("Name: %s\n", pokedexEntry.Name)
		fmt.Printf("Weight: %d\n", pokedexEntry.Weight)
		fmt.Printf("Height: %d\n", pokedexEntry.Height) 
		fmt.Println("Stats:")
		for _, value := range pokedexEntry.Stats {
			fmt.Printf("    -%s: %d\n", value.Stat.Name, value.BaseStat)
		}	
		fmt.Println("Types:")
		for _, value := range pokedexEntry.Types {
			fmt.Printf("    -Slot%d: %s\n", value.Slot, value.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}
