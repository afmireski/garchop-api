package entities

type Pokemon struct {
	Id string `json:"id"`
	ReferenceId uint `json:"reference_id"`
	Name string `json:"name"`
	Weight uint `json:"weight"`
	Height uint `json:"height"`
	ImageUrl string `json:"image_url"`
	Experience uint `json:"experience"`
	Types []PokemonType `json:"types"`
	Tier Tier `json:"tier"`
}

type PokemonProduct struct {
	Pokemon
	Price uint `json:"price"`
	InStock uint `json:"in_stock"`
}
