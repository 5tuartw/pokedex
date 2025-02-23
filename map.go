package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CommandMap(commands map[string]CliCommand, cfg *Config, param string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != nil {
		url = *cfg.Next
	}

	var locationResp LocationAreaResponse

	if data, ok := cfg.cache.Get(url); ok {
		//fmt.Println("Using cached data")
		err := json.Unmarshal(data, &locationResp)
		if err != nil{
			fmt.Printf("Error unmarshalling json location data from cache: %v", err)
			return err
		}
	} else {
		//fmt.Println("Requesting location data")
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
		cfg.cache.Add(url, data)

		err = json.Unmarshal(data, &locationResp)
		if err != nil {
			fmt.Printf("Error unmarshalling json location data: %v", err)
			return err
		}
	}

	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous
	cfg.Results = locationResp.Results
	cfg.Count = locationResp.Count

	for _, location := range cfg.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func CommandMapb(commands map[string]CliCommand, cfg *Config, param string) error {
	if cfg.Previous != nil {
		cfg.Next = cfg.Previous
		err := CommandMap(commands, cfg, param)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Already on first page")
	}
	return nil
}