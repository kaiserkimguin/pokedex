package main

import (
	"fmt"
)


func commandMap(cfg *configURL, args []string) error {
	locResp, err := cfg.pokeapiClient.ListLocations(cfg.next)
	if err != nil {
		return err
	}
	cfg.next = locResp.Next
	cfg.previous = locResp.Previous

	for _, value := range locResp.Results {
		fmt.Println(value.Name)
	}
	return nil
}

func commandMapb(cfg *configURL, args []string) error {
	if nil == cfg.previous {
		fmt.Println("you're on the first page")
		return nil
	}
	locResp, err := cfg.pokeapiClient.ListLocations(cfg.previous)
	if err != nil {
		return err
	}
	cfg.next = locResp.Next
	cfg.previous = locResp.Previous

	for _, value := range locResp.Results {
		fmt.Println(value.Name)
	}
	return nil
}
