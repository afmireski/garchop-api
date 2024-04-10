package models

import "time"

type PokemonModel struct {
	Id          string         `json:"id"`
	ReferenceId uint           `json:"reference_id"`
	TierId      string         `json:"tier_id"`
	Name        string         `json:"name"`
	Weight      uint            `json:"weight"`
	Height      uint            `json:"height"`
	ImageUrl    string         `json:"image_url"`
	Experience  uint            `json:"experience"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   *time.Time     `json:"deleted_at,omitempty"`
	Tier       *TierModel     `json:"tiers"`
	Types       []TypeModel `json:"types"`
	Prices      []PriceModel   `json:"prices"`
	Stock      *StockModel   `json:"stocks"`
}
