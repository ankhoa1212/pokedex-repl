package main

import "fmt"
import "strings"

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	return words
}