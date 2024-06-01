package pokemon

import (
	"strconv"
)

var (
	ErrNotFound = &PokemonNotFoundError{}
)

type PokemonNotFoundError struct {
	ID int `json:"id"`
}

func (e PokemonNotFoundError) Error() string {
	return "pokemon not found: " + strconv.Itoa(e.ID)
}

func ErrPokemonNotFound(id int) error {
	return PokemonNotFoundError{ID: id}
}
