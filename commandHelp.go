package main

import (
	"fmt"
)

func HelpCommand(c *config, cmds ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage")
	fmt.Println()
	for k, v := range getCommands() {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	fmt.Println()
	return nil
}
