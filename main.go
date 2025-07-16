package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Next             string
	Previous         string
}

type Command struct {
	name        string
	description string
	callback     func() error
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

func main() {
	var input string
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input = reader.Text()
		words := cleanInput(input)

		if cmd, ok := getCommands()[words[0]]; ok {
			cmd.callback()
		} else {
			fmt.Printf("Unknown command: %s\n", words[0])
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	return words
}

func commandMap() error {
	locationAreaEndpoint := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=%d", 0, 20)
	res, err := http.Get(locationAreaEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex...")
	os.Exit(0)
	return nil
}