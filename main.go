package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Commands    map[string]cliCommand
	NextUrl     string
	PreviousUrl string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		if err := reader.Err(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := cfg.Commands[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func main() {
	config := config{
		NextUrl:     "https://pokeapi.co/api/v2/location-area/",
		PreviousUrl: "",
		Commands: map[string]cliCommand{
			"help": {
				name:        "help",
				description: "Displays a help message",
				callback:    commandHelp,
			},
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"map": {
				name:        "map",
				description: "Displays the names of 20 location areas in the Pokemon world",
				callback:    commandMap,
			},
			"mapb": {
				name:        "map-back",
				description: "Displays the names of the previous 20 location areas in the Pokemon world",
				callback:    commandMapb,
			},
		},
	}
	startRepl(&config)
}
