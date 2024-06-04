package structs

type PokemonPickup struct {
	Pickup []struct {
		No string `json:"no"`
	} `json:"pickup"`
}
