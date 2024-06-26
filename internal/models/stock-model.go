package models

import "time"

type StockModel struct {
	PokemonId string     `json:"pokemon_id"`
	Quantity  uint        `json:"quantity"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Pokemons  *PokemonModel
}
