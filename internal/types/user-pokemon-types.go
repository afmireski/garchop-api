package types

type UserPokemonData struct {
	UserId    string `json:"user_id"`
	PokemonId string `json:"pokemon_id"`
	Quantity  uint   `json:"quantity"`
}

type UpdateUserPokemonData struct {
	Quantity uint `json:"quantity"`
}
