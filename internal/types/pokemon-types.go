package types

type CreatePokemonInput struct {
	ReferenceId int    `json:"reference_id"`
	TierId      string `json:"tier_id"`
	Name        string `json:"name"`
	Weight      int    `json:"weight"`
	Height      int    `json:"height"`
	ImageUrl    string `json:"image_url"`
	Experience  int    `json:"experience"`
}

type NewPokemonInput struct {
	Name         string `json:"name"`
	Price        int    `json:"price"`
	InitialStock int    `json:"stock"`
	TierId       string `json:"tier_id"`
}
