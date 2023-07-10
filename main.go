package main

import (
	"fmt"

	"github.com/polyglotdev/pokeapi/pokeapi"
)

func main() {
	client := pokeapi.NewClient()
	pokemon, err := client.GetPokemon("pikachu")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Pokemon:", pokemon.Name)

	nature, err := client.GetNature("adamant")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Nature:", nature.Name)
}
