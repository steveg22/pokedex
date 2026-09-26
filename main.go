package main

import (
	"time"

	"github.com/steveg22/pokedexcli/internal/pokeapi"
	"github.com/steveg22/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	pokeCache := pokecache.NewCache(time.Second * 5)
	pokeCache.Add("result", []byte("victory"))
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}
