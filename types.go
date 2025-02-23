package main

import (
	"pokedex/internal/pokecache"
)
type CliCommand struct {
	name        string
	description string
	callback    func(map[string]CliCommand, *Config, string) error
}

type LocationInfo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Config struct {
	cache 			*pokecache.Cache
	Count		    int
	Next		    *string
	Previous		*string
	Results  		[]LocationInfo
	CaughtPokemon	[]Pokemon
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationInfo `json:"results"`
}

type LocalPokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonEncounter struct {
	LocalPokemon    LocalPokemon `json:"pokemon"`
}

type LocationData struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name 		string			`json:"name"`
	Experience	int				`json:"base_experience"`
	Height		int				`json:"height"`
	Weight		int				`json:"weight"`
	Stats		[]PokemonStat	`json:"stats"`
	Types		[]PokemonType	`json:"types"`
}

type PokemonStat struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     struct {
		Name string `json:"name"`
	}
}

type PokemonType struct {
	Type struct {
		Name string `json:"name"`
	}
}