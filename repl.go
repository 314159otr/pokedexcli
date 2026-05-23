package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		clean := cleanInput(input)
		firstWord := ""
		if len(clean) > 0 {
			firstWord = clean[0]
		}
		fmt.Println("Your command was: " + firstWord)
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)

	return words
}
