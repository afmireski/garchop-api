package types

type CreatePokemonInput struct {
	ReferenceId int `json:"reference_id"`
	TierId string `json:"tier_id"`
	Name string `json:"name"`
	Weight int `json:"weight"`
	Height int `json:"height"`
	ImageUrl string `json:"image_url"`
	Experience int `json:"experience"`
}