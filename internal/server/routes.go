package server

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"ccpokedex/cmd/web"
	pw "ccpokedex/cmd/web/pokemon"
	p "ccpokedex/internal/server/pokemon/handler"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/api", jsonMarshal(s.helloWorldHandler))
	r.Get("/api/health", jsonMarshal(s.healthHandler))

	pokemon := p.NewPokemonRouteHandler(s.pokemon)
	r.Get("/api/pokemon/trends", jsonMarshal(pokemon.HandleGetPokemonTrends))
	r.Get("/api/pokemon/{id}", jsonMarshal(pokemon.HandleGetPokemon))

	r.Handle("/assets/*", web.Public())
	r.Get("/web", templ.Handler(web.HelloForm()).ServeHTTP)
	r.Post("/hello", web.HelloWebHandler)
	r.Get("/pokemon", templ.Handler(pw.PokedexIndex()).ServeHTTP)
	r.Get("/pokemon/{id}", pokemon.HandleGetPokemonDetail)

	r.Get("/web/pokemon/trends", pokemon.HandleGetHtmxPokemonTrends)
	r.Get("/web/pokemon/{id}", pokemon.HandleGetHtmxPokemon)

	return r
}

type RestHandler func(w http.ResponseWriter, r *http.Request) (interface{}, error)

func jsonMarshal(handler RestHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload, err := handler(w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp, err := json.Marshal(payload)
		if err != nil {
			slog.Error("error handling JSON marshal", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if resp == nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		_, _ = w.Write(resp)
	})
}

func (s *Server) helloWorldHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"
	return resp, nil
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return s.db.Health(), nil
}
