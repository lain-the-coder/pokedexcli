package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	cleanedSlice := strings.Fields(strings.ToLower(text))
	return cleanedSlice
}

func getCommands() map[string]cliCommand {
	// 1. Create an empty map first
	commands := make(map[string]cliCommand)

	// 2. Add the exit command to it manually
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	// 3. Add the help command
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback: func() error {
			fmt.Println("\nWelcome to the Pokedex!")
			fmt.Println("Usage:")

			// Because this function is written right here,
			// it can see and loop over the 'commands' map perfectly!
			for _, cmd := range commands {
				fmt.Printf("%s: %s\n", cmd.name, cmd.description)
			}
			fmt.Println()
			return nil
		},
	}

	return commands
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleaned := cleanInput(scanner.Text())

		if len(cleaned) == 0 {
			continue
		}

		// 2. Grab the first word they typed
		commandName := cleaned[0]

		// 3. Look up that word in our commands map
		command, exists := commands[commandName]

		if exists {
			// 4. If it exists, execute its callback function!
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// 5. If it doesn't exist, print the unknown message
			fmt.Println("Unknown command")
		}
	}
}
