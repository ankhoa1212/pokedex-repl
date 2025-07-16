package main

import "fmt"
import "strings"
import "bufio"
import "os"

func main() {
	var input string
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input = reader.Text()
		words := cleanInput(input)
		for i := range words {
			if words[i] == "help" {
				commandHelp()
			} else if words[i] == "exit" {
				commandExit()
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