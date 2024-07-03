package entities

import "github.com/afmireski/garchop-api/internal/types/enums"

type Reward struct {
	Id                 string              `json:"id"`
	TierId             uint                `json:"tier_id"`
	Name               string              `json:"name"`
	Description        *string             `json:"description"`
	ExperienceRequired uint                `json:"experience_required"`
	Prize              string              `json:"prize"`
	PrizeType          enums.PrizeTypeEnum `json:"prize_type"`
	Tier               *Tier               `json:"tier"`
}
