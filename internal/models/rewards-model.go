package models

import (
	"time"

	"github.com/afmireski/garchop-api/internal/types/enums"

	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type RewardModel struct {
	Id                 string              `json:"id"`
	TierId             uint                `json:"tier_id"`
	Name               string              `json:"name"`
	Description        *string             `json:"description"`
	ExperienceRequired uint                `json:"experience_required"`
	Prize              myTypes.AnyMap      `json:"prize"`
	PrizeType          enums.PrizeTypeEnum `json:"prize_type"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
	DeletedAt          *time.Time          `json:"deleted_at,omitempty"`
	Tier               *TierModel          `json:"tiers"`
}
