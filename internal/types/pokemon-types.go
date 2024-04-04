package types

type CreatePokemonInput struct {
	ReferenceId uint    `json:"reference_id"`
	TierId      string `json:"tier_id"`
	Name        string `json:"name"`
	Weight      uint    `json:"weight"`
	Height      uint    `json:"height"`
	ImageUrl    string `json:"image_url"`
	Experience  uint    `json:"experience"`
}

type NewPokemonInput struct {
	Name         string `json:"name"`
	Price        int    `json:"price"`
	InitialStock int    `json:"initial_stock"`
	TierId       string `json:"tier_id"`
}
