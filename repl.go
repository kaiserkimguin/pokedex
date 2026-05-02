package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/kaiserkimguin/pokedex/internal/pokeapi"

)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &configURL{
		pokeapiClient: pokeapi.NewClient(5 * time.Second, 5 * time.Second),
		pokedex: make(map[string]pokeapi.FullPokemon),
	}
	for {
		scanner.Scan()
		prompt := scanner.Text()
		promptWords := cleanInput(prompt)
		cmdMap := getCommands()
		command, exists := cmdMap[promptWords[0]]
		if exists {
			fmt.Println("Pokedex >", command.name)
			err := command.callback(cfg, promptWords[1:]); if err != nil {
				fmt.Println("Command not found")
			}
			}
		}
}

type configURL struct {
	next     				*string
	previous	 			*string
	pokeapiClient 	pokeapi.Client
	pokedex 				map[string]pokeapi.FullPokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *configURL, args []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		// Map Funktion, holt die nächsten 20 Werte von der Internetseite
		"map": {
			name:        "map",
			description: "displays the names of the next 20 locations in the Pokemon world",
			callback:    commandMap,
		},
		// Map backwards Funktion
		// diese Funktion zieht die letzten 20 Orte. Dafür holt sie alle Werte der letzten Seite
		"mapb": {
			name:        "mapb",
			description: "displays the names of the previous 20 locations in the Pokemon world",
			callback:    commandMapb,
		},
		// explore functions. Mit einem Ort als argv gibt sie alle fangbaren
		// pokemon zurueck
		"explore": {
			name: 				"explore",
			description:	"displays all possible pokemon encounters of and area",
			callback: 		commandExplore,
		},
		// catch function, enables user to catch pokemon. Gives him a chacne 
		// to catch a pokemon based on its experience
		"catch": {
			name: 				"catch",
			description: 	"allows user to add pokemon to his pokedex",
			callback: 		commandCatch,
		},	
	}
}

func cleanInput(text string) []string {
	// lowering the text
	text = strings.ToLower(text)
	// trimming the string from whitespaces
	text = strings.TrimSpace(text)
	// dividing the string
	result := strings.Fields(text)
	return result
}
