package types

import (
	"encoding/json"

	"github.com/afmireski/garchop-api/internal/types/enums"
)


type CreateRewardInput struct {
	TierId             uint                `json:"tier_id"`
	Name               string              `json:"name"`
	Description        string              `json:"description"`
	ExperienceRequired uint                `json:"experience_required"`
	Type               enums.PrizeTypeEnum `json:"prize_type"`
	Prize              json.RawMessage     `json:"prize"`
}
