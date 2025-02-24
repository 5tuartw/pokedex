package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokecache"
	"strings"
	"time"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	cfg := &Config{
		cache: pokecache.NewCache(5 * time.Minute),
	}
	for {
		if scanner.Scan() {
			input := scanner.Text()
			inputclean := cleanInput(input)
			input = inputclean[0]
			param := ""
			if len(inputclean) > 1 {
				param = inputclean[1]
			}
			if command, ok := commands[input]; ok {
				command.callback(commands, cfg, param)
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

var commands = map[string]CliCommand{
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
		callback:    CommandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "list the previous 20 locations",
		callback:    CommandMapb,
	},
	"explore": {
		name:        "explore <location>",
		description: "list the encounters for a location",
		callback:    CommandExplore,
	},
	"catch": {
		name:        "catch <Pokemon>",
		description: "attempt to catch a Pokemon",
		callback:    CommandCatch,
	},
	"inspect": {
		name:        "inspect <Pokemon>",
		description: "see the details of a caught Pokemon",
		callback:    CommandInspect,
	},
	"dex": {
		name:        "dex",
		description: "see the pokemon in your pokedex",
		callback:    CommandPokedex,
	},
}

func commandExit(commands map[string]CliCommand, cfg *Config, param string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands map[string]CliCommand, cfg *Config, param string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for c := range commands {
		fmt.Printf("%v: %v\n", commands[c].name, commands[c].description)
	}
	return nil
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
