package types

type CreatePokemonInput struct {
	ReferenceId uint   `json:"reference_id"`
	TierId      int    `json:"tier_id"`
	Name        string `json:"name"`
	Weight      uint   `json:"weight"`
	Height      uint   `json:"height"`
	ImageUrl    string `json:"image_url"`
	Experience  uint   `json:"experience"`
}

type RegistryPokemonInput struct {
	CreatePokemonInput
	Price        int      `json:"price"`
	InitialStock int      `json:"initial_stock"`
	Types        []string `json:"types"`
}

type NewPokemonInput struct {
	Name         string `json:"name"`
	Price        int    `json:"price"`
	InitialStock int    `json:"initial_stock"`
	TierId       int    `json:"tier_id"`
}

type CreatePokemonTypeInput struct {
	ReferenceId uint64 `json:"reference_id"`
	Name        string `json:"name"`
}
