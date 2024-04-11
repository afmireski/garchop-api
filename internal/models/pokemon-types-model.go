package models

type PokemonTypesModel struct {
	PokemonId string `json:"pokemon_id"`
	TypeId string `json:"type_id"`
	Pokemons *PokemonModel `json:"pokemons"`
	Types *TypeModel `json:"types"`
}