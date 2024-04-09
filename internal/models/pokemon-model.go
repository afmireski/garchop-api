package models

import "time"

type PokemonModel struct {
	id string `json:"id"`
	ReferenceId uint `json:"reference_id"`
	TierId string `json:"tier_id"`
	Name string `json:"name"`
	Weight int `json:"weight"`
	Height int `json:"height"`
	ImageUrl string `json:"image_url"`
	Experience int `json:"experience"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}