package main

import (
	"fmt"
	"os"
)

func commandExit(c *configURL) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
