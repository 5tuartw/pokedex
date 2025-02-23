package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	cfg := &Config{}
	for {
		if scanner.Scan() {
			input := scanner.Text()
			inputclean := cleanInput(input)
			input = inputclean[0]
			if command, ok := commands[input]; ok {
				command.callback(commands, cfg)
			} else {
				fmt.Println("Uknown command")
			}
			//fmt.Printf("Your command was: %v\n", inputclean[0])
			fmt.Print("Pokedex > ")

		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading command: ", err)
		}
	}
}

func commandExit(commands map[string]cliCommand, cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands map[string]cliCommand, cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for c := range commands {
		fmt.Printf("%v: %v\n", commands[c].name, commands[c].description)
	}
	return nil
}

func commandMap(commands map[string]cliCommand, cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error requesting locations: %v", err)
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading location data: %v", err)
		return err
	}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Printf("Error unmarshalling json location data: %v", err)
		return err
	}
	for _, location := range cfg.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(commands map[string]cliCommand, cfg *Config) error {
	if cfg.Previous != nil {
		cfg.Next = cfg.Previous
		err := commandMap(commands, cfg)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Already on first page")
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(map[string]cliCommand, *Config) error
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Lists the next 20 locations",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "list the previous 20 locations",
		callback:    commandMapb,
	},
}

type LocationInfo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Config struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationInfo `json:"results"`
}

func cleanInput(text string) []string {
	var output []string
	ltext := strings.ToLower(text)
	//fmt.Printf("Lowering case: %v\n", ltext)
	words := strings.Split(ltext, " ")
	//fmt.Printf("Splitting by spaces: %v\n", words)
	for i := range words {
		if words[i] != "" {
			output = append(output, strings.TrimSpace(words[i]))
		}
	}
	//fmt.Printf("Trimmed spaces: %v\n", output)
	return output
}
