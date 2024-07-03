package models

import (
	"time"
)

type PrizeTypeEnum string

const (
	Pokemon PrizeTypeEnum = "pokemon"
)

type RewardModel struct {
	Id                 string        `json:"id"`
	TierId             uint          `json:"tier_id"`
	Name               string        `json:"name"`
	Description        *string       `json:"description"`
	ExperienceRequired uint          `json:"experience_required"`
	Prize              string        `json:"prize"`
	PrizeType          PrizeTypeEnum `json:"prize_type"`
	CreatedAt          time.Time     `json:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at"`
	DeletedAt          *time.Time    `json:"deleted_at,omitempty"`
}
