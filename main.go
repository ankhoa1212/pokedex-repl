package main

import "fmt"
import "strings"
import "bufio"
import "os"

type Command struct {
	name        string
	description string
	callback     func() error
}

map[string]Command{
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

func main() {
	var input string
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input = reader.Text()
		words := cleanInput(input)
		for i := range words {
			if cmd, ok := commands[words[i]]; ok {
				cmd.callback()
			}
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	return words
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	fmt.Println("  help: Show this help message")
	fmt.Println("  exit: Exit the Pokedex")
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex...")
	os.Exit(0)
	return nil
}