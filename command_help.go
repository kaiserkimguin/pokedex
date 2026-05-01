package main

import (
	"fmt"
)

func commandHelp(c *configURL) error {
	cmdMap := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage: \n\n\n")
	for key, value := range cmdMap {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}
