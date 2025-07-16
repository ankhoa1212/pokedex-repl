package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	name        string
	description string
	callback     func() error
}

func startPokedex() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())

		if command, found := getCommands()[words[0]]; found {
			err := command.callback()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", words[0])
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	return words
}

func getCommands() map[string]Command {
	return map[string]Command{
		"map": {
			name:        "map",
			description: "Show the names of 20 location areas. Each subsequent call will show the next 20 location areas.",
			callback:    commandMap,
		},
		"help": {
			name:        "help",
			description: "Show the help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}