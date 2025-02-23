package main

import (
	"fmt"
	"strings"
	"net/http"
	"io"
	"encoding/json"
)

func CommandExplore(commands map[string]CliCommand, cfg *Config, param string) error {
	url := "https://pokeapi.co/api/v2/location-area/"+param+"/"
	var locationData LocationData
	if data, ok := cfg.cache.Get(url); ok {
		//fmt.Println("Using cached encounter data")
		err := json.Unmarshal(data, &locationData)
		if err != nil{
			fmt.Printf("Error unmarshalling json encounter data from cache: %v", err)
			return err
		}
	} else {
		//fmt.Println("Requesting encounter data")
		res, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error requesting encounters: %v", err)
			return err
		}
		defer res.Body.Close()
        
		if res.StatusCode == http.StatusNotFound {
            fmt.Println("Location not found")
            return nil
        }

		data, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Error reading encounter data: %v", err)
			return err
		}
		cfg.cache.Add(url, data)
		err = json.Unmarshal(data, &locationData)
		if err != nil {
			fmt.Printf("Error unmarshalling json encounter data: %v", err)
			return err
		}
	}

	fmt.Printf("Found the following %v pokemon at %v:\n", len(locationData.PokemonEncounters), param)
	for _, encounter := range locationData.PokemonEncounters {
		fmt.Println(strings.Title(encounter.LocalPokemon.Name))
	}

	return nil
}