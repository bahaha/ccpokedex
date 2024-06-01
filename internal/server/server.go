package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"ccpokedex/internal/database"
	"ccpokedex/internal/server/pokemon"
	pokeapi "ccpokedex/internal/server/pokemon/poke_api"
)

type Server struct {
	port int

	db database.Service

	pokemon pokemon.PokemonRepo
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,

		db:      database.New(),
		pokemon: pokeapi.NewPokeAPI(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
