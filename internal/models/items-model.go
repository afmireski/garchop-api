package models

import "time"

type ItemModel struct {
	Id         string        `json:"id"`
	CartId     *string        `json:"cart_id"`
	PokemonId  string        `json:"pokemon_id"`
	PurchaseId *string       `json:"purchase_id",omitempty`
	Quantity   uint          `json:"quantity"`
	Price      uint          `json:"price"`
	Total      uint          `json:"total"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
	DeletedAt  *time.Time    `json:"deleted_at,omitempty"`
	Pokemon    *PokemonModel `json:"pokemons"`
}
