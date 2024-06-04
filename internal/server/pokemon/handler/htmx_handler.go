package handler

import (
	web "ccpokedex/cmd/web/pokemon"
	"log/slog"
	"net/http"
)

func (h *PokemonRouteHandler) HandleGetHtmxPokemon(
	w http.ResponseWriter,
	r *http.Request,
) {
	pokemon, err := h.handleGetPokemon(w, r)
	if err != nil {
		slog.Error("failed to get pokemon", "err", err)
		return
	}

	component := web.PokemonProfile(pokemon)
	err = component.Render(r.Context(), w)

	if err != nil {
		slog.Error("failed to render pokemon profile", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *PokemonRouteHandler) HandleGetHtmxPokemonTrends(
	w http.ResponseWriter,
	r *http.Request,
) {
	pokemons, err := h.handleGetPokemonTrends(w, r)
	if err != nil {
		slog.Error("failed to get pokemon trends", "err", err)
		return
	}

	component := web.PokemonTrends(pokemons)
	err = component.Render(r.Context(), w)

	if err != nil {
		slog.Error("failed to render pokemon trends", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
