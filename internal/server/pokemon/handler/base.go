package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	p "ccpokedex/internal/server/pokemon"
)

type PokemonRouteHandler struct {
	repo p.PokemonRepo
}

func NewPokemonRouteHandler(repo p.PokemonRepo) *PokemonRouteHandler {
	return &PokemonRouteHandler{
		repo: repo,
	}
}

func (h *PokemonRouteHandler) handleGetPokemon(
	w http.ResponseWriter,
	r *http.Request,
) (*p.Pokemon, error) {
	id := chi.URLParam(r, "id")
	pokeId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return nil, err

	}

	lang := r.URL.Query().Get("lang")
	if lang == "" {
		lang = "en"
	}

	return h.repo.GetPokemon(pokeId, lang)
}
