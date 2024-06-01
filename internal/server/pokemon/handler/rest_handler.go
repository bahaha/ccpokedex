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
