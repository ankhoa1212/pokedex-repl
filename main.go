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
		fmt.Println("You entered:", words)
		if words[0] == "exit" {
			commandExit()
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	return words
}

func commandExit() error {
	fmt.Println("Closing the Pokedex...")
	os.Exit(0)
	return nil
}