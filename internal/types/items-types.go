package types

type CreateItemInput struct {
	CartId    string `json:"cart_id"`
	PokemonId string `json:"pokemon_id"`
	Price     uint   `json:"price"`
	Quantity  uint   `json:"quantity"`
	Total     uint   `json:"total"`
}
