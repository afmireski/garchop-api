package models

type PokemonTypes struct {
	PokemonId string `json:"pokemon_id"`
	TypeId string `json:"type_id"`
	Pokemons *PokemonModel
	Types *TypeModel
}