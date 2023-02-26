package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/henryckl/pokerequest/database"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

type Pokemon struct {
	Name   string `json:"name"`
	Height uint   `json:"height"`
	Weight uint   `json:"weight"`
	ID     uint   `json:"id"`
}

func getPokemon(url string) (pokemon *Pokemon, err error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(response.Body).Decode(&pokemon)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

func insertPokemon(pokemons []Pokemon) {
	db := database.Init()
	for _, pokemon := range pokemons {
		db.Exec("inserirPokemon", pokemon.Name, pokemon.Height, pokemon.Weight, pokemon.ID)
	}
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Execution Time: ", time.Since(start))
	}()
	var urls []string
	var pokemons []Pokemon
	for _, value := range makeRange(1, 1000) {
		urls = append(urls, fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", value))
	}
	ch := make(chan Pokemon)
	for _, url := range urls {
		go func(url string) {
			poke, err := getPokemon(url)
			if err != nil {
				return
			}
			ch <- *poke
		}(url)
	}
	for i := 0; i < len(makeRange(1, 1000)); i++ {
		pokemon := <-ch
		pokemons = append(pokemons, pokemon)
	}
	insertPokemon(pokemons)
	fmt.Printf("Total carregado: %d \n", len(pokemons))
}
