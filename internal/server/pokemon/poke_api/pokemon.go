package pokeapi

import (
	p "ccpokedex/internal/server/pokemon"
	"ccpokedex/internal/server/pokemon/poke_api/structs"
	"fmt"
	"log/slog"
	"net/http"
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

	types := []string{}
	for _, t := range pokemon.Types {
		types = append(types, t.Type.Name)
	}

	return &p.Pokemon{
		ID:     pokemon.ID,
		Name:   pokemon.Name,
		Avatar: pokemon.Sprites.Other.OfficialArtwork.FrontDefault,
		Types:  types,
	}, nil
}
