package main

import (
	"strings"
)

func cleanInput(text string) []string {
	cleanedSlice := strings.Fields(strings.ToLower(text))
	return cleanedSlice
}
func main() {
	cleanInput("Charmander Bulbasaur PIKACHU")
}
