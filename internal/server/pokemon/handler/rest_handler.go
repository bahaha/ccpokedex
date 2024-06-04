package handler

import (
	"errors"
	"net/http"

	p "ccpokedex/internal/server/pokemon"
)

func (h *PokemonRouteHandler) HandleGetPokemon(
	w http.ResponseWriter,
	r *http.Request,
) (interface{}, error) {
	pokemon, err := h.handleGetPokemon(w, r)
	if errors.As(err, p.ErrNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return nil, err
	}
	return pokemon, err
}

func (h *PokemonRouteHandler) HandleGetPokemonTrends(
	w http.ResponseWriter,
	r *http.Request,
) (interface{}, error) {
	pokemons, err := h.handleGetPokemonTrends(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	return pokemons, nil
}
