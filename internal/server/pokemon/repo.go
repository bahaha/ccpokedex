package pokemon

type PokemonType struct {
	Name  string `json:"name"`
	Badge string `json:"badge"`
}

type Pokemon struct {
	ID     int           `json:"id"`
	Name   string        `json:"name"`
	Title  string        `json:"title"`
	Avatar string        `json:"avatar"`
	Types  []PokemonType `json:"types"`
}

func NewPokemon() *Pokemon {
	return &Pokemon{}
}

type PokemonRepo interface {
	GetPokemonTrends(lang string) ([]*Pokemon, error)
	GetPokemon(id int, lang string) (*Pokemon, error)
}
