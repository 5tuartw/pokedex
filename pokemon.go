package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
	"math/rand"
)

func CommandCatch(commands map[string]CliCommand, cfg *Config, name string) error {
	url := "https://pokeapi.co/api/v2/pokemon/"+name+"/"
	var pokemon Pokemon
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error requesting pokemon data: %v", err)
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading pokemon data: %v", err)
		return err
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		fmt.Printf("Error unmarshalling json pokemon data: %v", err)
		return err
	}
	//fmt.Printf("Caught %v with base experience %v\n", pokemonData.Name, pokemonData.Experience)
	fmt.Printf("Throwing a Pokeball at %v...\n", name)
    // Calculate the initial chance to catch the Pokémon
    baseExperience := float64(pokemon.Experience)
    maxBaseExperience := 1000.0
    initialChance := 1.0 - (baseExperience / maxBaseExperience) - 0.5
    // Ensure the initial chance is within a reasonable range
    if initialChance < 0.1 {
        initialChance = 0.1
    }
    //fmt.Printf("Initial chance of catching: %.2f%%\n", initialChance*100)
    // Player has 3 attempts to catch the Pokémon
    for attempt := 1; attempt <= 3; attempt++ {
		fmt.Printf("Attempt %d: Throwing ball", attempt)
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
		fmt.Printf(".\n")
        roll := rand.Float64()

        if roll < initialChance {
            fmt.Printf("Caught %v on attempt %d!\n", name, attempt)
			cfg.CaughtPokemon = append(cfg.CaughtPokemon, pokemon)
            return nil
        }
        // Increase the chance slightly for the next attempt
        initialChance += 0.1
        if initialChance > 1.0 {
            initialChance = 1.0
        }
		fmt.Println("Missed!")
    }

    //fmt.Printf("Failed to catch %v after 3 attempts.\n", name)
	return nil
}

func CommandInspect(commands map[string]CliCommand, cfg *Config, name string) error {
	for _, pokemon := range cfg.CaughtPokemon {
		if pokemon.Name == name {
			fmt.Printf("Name: \033[32m%v\033[0m\n", strings.Title(pokemon.Name))
			fmt.Printf("Height: \033[32m%v\033[0m\n", pokemon.Height)
			fmt.Printf("Weight: \033[32m%v\033[0m\n", pokemon.Weight)
			fmt.Println("Stats:")
			for _, stat := range pokemon.Stats {
				fmt.Printf("  -%v: \033[32m%v\033[0m\n", stat.Stat.Name, stat.BaseStat)
			}
			fmt.Println("Types:")
			for _, typeData := range pokemon.Types {
				fmt.Printf("  - \033[32m%v\033[0m\n", typeData.Type.Name)
			}
			return nil
		}
	}
	fmt.Printf("You have not caught %v\n", name)
	return nil
}

func CommandPokedex(commands map[string]CliCommand, cfg *Config, name string) error {

	if cfg.CaughtPokemon == nil {
		fmt.Println("You have not caught any Pokemon")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.CaughtPokemon {
		fmt.Printf(" - \033[34m%v\033[0m\n", strings.Title(pokemon.Name))
	}	

	return nil
}