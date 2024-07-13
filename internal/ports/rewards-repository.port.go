package ports

import (
	"encoding/json"

	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
	"github.com/afmireski/garchop-api/internal/types/enums"
)

type RewardsRepositoryPort interface {
	Create(input CreateRewardInput) (string, error)
	FindAll(where myTypes.Where) ([]models.RewardModel, error)
	Delete(id string, where myTypes.Where) error
}

type CreateRewardInput struct {
	TierId             uint                `json:"tier_id"`
	Name               string              `json:"name"`
	Description        string              `json:"description"`
	ExperienceRequired uint                `json:"experience_required"`
	Type               enums.PrizeTypeEnum `json:"prize_type"`
	Prize              json.RawMessage     `json:"prize"`
}
