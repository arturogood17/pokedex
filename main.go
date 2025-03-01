package main

import (
	"strings"
	"time"

	"github.com/arturogood17/pokedex/internal/pokeapi"
)

func cleanInput(text string) []string {
	if len(text) <= 0 {
		return []string{}
	}
	words := strings.Fields(strings.ToLower(text))
	return words
}

func main() {

	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeClient: pokeClient,
	}
	startREPL(cfg)
}
