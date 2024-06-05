package pokeapi

import (
	p "ccpokedex/internal/server/pokemon"
	"ccpokedex/internal/server/pokemon/poke_api/structs"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var (
	base = "https://pokeapi.co/api/v2"
)

type PokeAPI struct {
}

func NewPokeAPI() *PokeAPI {
	return &PokeAPI{}
}

func (api *PokeAPI) GetPokemon(id int, lang string) (*p.Pokemon, error) {
	pokemon := &structs.Pokemon{}

	status, err := send(fmt.Sprintf("/pokemon/%d", id), pokemon)
	if status == http.StatusNotFound {
		return nil, p.ErrPokemonNotFound(id)
	}

	if err != nil {
		slog.Error("error getting pokemon", "error", err)
		return nil, err
	}

	types := []p.PokemonType{}
	for _, t := range pokemon.Types {
		typ := &p.PokemonType{
			Name:  t.Type.Name,
			Badge: fmt.Sprintf("/assets/pokemon/types/%s.png", strings.ToLower(t.Type.Name)),
		}
		types = append(types, *typ)
	}

	return &p.Pokemon{
		ID:     pokemon.ID,
		Name:   pokemon.Name,
		Avatar: pokemon.Sprites.Other.OfficialArtwork.FrontDefault,
		Types:  types,
	}, nil
}

func (api *PokeAPI) GetPokemonTrends(lang string) ([]*p.Pokemon, error) {
	pokemons := []*p.Pokemon{}

	trends := &structs.PokemonPickup{}
	_, err := send("/pickup", trends)
	if err != nil {
		slog.Error("error getting pokemon trends", "error", err)
		return nil, err
	}

	wg := &sync.WaitGroup{}
	getPokemonChan := make(chan *p.Pokemon)
	for _, pokemon := range trends.Pickup {
		id, err := strconv.Atoi(pokemon.No)
		if err != nil {
			slog.Error("error parsing pokemon id", "error", err)
			continue
		}

		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// TODO: handle get pokemon error
			pokemon, _ := api.GetPokemon(id, lang)
			getPokemonChan <- pokemon
		}(id)
	}

	go func() {
		wg.Wait()
		close(getPokemonChan)
	}()

	for pokemon := range getPokemonChan {
		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}
