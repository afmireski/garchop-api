package models

import "time"

type UserStatsModel struct {
	UserId       string     `json:"user_id"`
	TierId       uint       `json:"tier_id"`
	Experience   uint       `json:"experience"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
	PreviousTier *TierModel `json:"tiers,omitempty"`
}
