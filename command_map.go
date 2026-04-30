package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(c *configURL) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.next != nil {
		url = *c.next
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("unable to make request")
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("unable to get response")
		return err
	}
	defer res.Body.Close()

	var loc location

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&loc); err != nil {
		fmt.Println("unable to read response")
		return err
	}
	c.next = loc.Next
	c.previous = loc.Previous

	for _, value := range loc.Results {
		fmt.Println(value.Name)
	}
	return nil
}

func commandMapb(c *configURL) error {
	var url string

	if c.previous != nil {
		url = *c.previous
	} else {
		fmt.Println("you're on the first page")
		return nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("unable to make request")
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("unable to get response")
		return err
	}
	defer res.Body.Close()

	var loc location

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&loc); err != nil {
		fmt.Println("unable to read response")
		return err
	}
	c.next = loc.Next
	c.previous = loc.Previous

	for _, value := range loc.Results {
		fmt.Println(value.Name)
	}
	return nil
}
