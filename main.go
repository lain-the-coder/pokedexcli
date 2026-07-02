package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2/location-area"

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	Next     string
	Previous string
	Client   *http.Client
}

type locationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []locationArea `json:"results"`
}

type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	cleanedSlice := strings.Fields(strings.ToLower(text))
	return cleanedSlice
}

func commandMap(cfg *config) error {
	currentURL := baseURL
	if cfg.Next != "" {
		currentURL = cfg.Next
	}
	err := doRequest(cfg, currentURL)
	if err != nil {
		return err
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	currentURL := cfg.Previous
	err := doRequest(cfg, currentURL)
	if err != nil {
		return err
	}
	return nil
}

func doRequest(cfg *config, currentURL string) error {
	u, err := url.Parse(currentURL)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}
	res, err := cfg.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("unexpected status from GET %s: %d %s", currentURL, res.StatusCode, res.Status)
	}
	var result locationAreaResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return err
	}
	for _, r := range result.Results {
		fmt.Println(r.Name)
	}
	if result.Next == nil {
		cfg.Next = ""
	} else {
		cfg.Next = *result.Next
	}
	if result.Previous == nil {
		cfg.Previous = ""
	} else {
		cfg.Previous = *result.Previous
	}
	return nil
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
		callback: func(_ *config) error {
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

	// 4. Add the map command
	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the next 20 location areas",
		callback:    commandMap,
	}

	// 5. Add the mapb command
	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the previous 20 location areas",
		callback:    commandMapb,
	}

	return commands
}

func main() {
	client := &http.Client{Timeout: 10 * time.Second}
	cfg := config{
		Next:     "",
		Previous: "",
		Client:   client,
	}
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
			err := command.callback(&cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// 5. If it doesn't exist, print the unknown message
			fmt.Println("Unknown command")
		}
	}
}
