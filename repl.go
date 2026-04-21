package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		prompt := scanner.Text()
		promptWords := cleanInput(prompt)
		cmdMap := getCommands()
		command, exists := cmdMap[promptWords[0]]
		if exists {
			command.callback()
		} else {
			fmt.Print("Command not found")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
		"map": {
			name:        "map",
			description: "displays the names of the next 20 locations in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the names of the previous 20 locations in the Pokemon world",
			callback:    commandMapb,
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
