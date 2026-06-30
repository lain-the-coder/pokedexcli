package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	cleanedSlice := strings.Fields(strings.ToLower(text))
	return cleanedSlice
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleaned := cleanInput(scanner.Text())
		if len(cleaned) > 0 {
			fmt.Printf("Your command was: %s\n", cleaned[0])
		}
	}
}
