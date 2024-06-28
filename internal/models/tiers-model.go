package models

import "time"

type TierModel struct {
	Id                uint       `json:"id"`
	PreviousTierId    uint       `json:"previous_tier_id"`
	Name              string     `json:"name"`
	MinimalExperience uint       `json:"minimal_experience"`
	LimitExperience   uint       `json:"limit_experience"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`
	PreviousTier      *TierModel `json:"previous_tier"`
}
