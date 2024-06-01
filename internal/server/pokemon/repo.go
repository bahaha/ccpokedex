package pokemon

type Pokemon struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Title  string   `json:"title"`
	Avatar string   `json:"avatar"`
	Types  []string `json:"types"`
}

func NewPokemon() *Pokemon {
	return &Pokemon{}
}

type PokemonRepo interface {
	GetPokemon(id int, lang string) (*Pokemon, error)
}
