package models

import "time"

type PriceModel struct {
	PokemonId string `json:"pokemon_id"`
	CreatedAt time.Time `json:"created_at"`
	Value int `json:"value"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Pokemons *PokemonModel
}
