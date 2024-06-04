package handler

import (
	"log/slog"
	"net/http"

	web "ccpokedex/cmd/web/pokemon"
)

func (h *PokemonRouteHandler) HandleGetPokemonDetail(
	w http.ResponseWriter,
	r *http.Request,
) {
	pokemon, err := h.handleGetPokemon(w, r)
	if err != nil {
		slog.Error("failed to get pokemon", "err", err)
		return
	}

	component := web.PokemonDetail(pokemon)
	err = component.Render(r.Context(), w)

	if err != nil {
		slog.Error("failed to render pokemon profile", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
