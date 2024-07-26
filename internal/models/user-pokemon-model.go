package models

type UserPokemonModel struct {
	UserId    string `json:"user_id"`
	PokemonId string `json:"pokemon_id"`
	Quantity  uint   `json:"quantity"`
}
