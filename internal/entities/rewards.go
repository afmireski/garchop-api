package entities

import (
	"github.com/afmireski/garchop-api/internal/models"
	"github.com/afmireski/garchop-api/internal/types/enums"
)

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

func BuildRewardFromModel(model *models.RewardModel) *Reward {
	return &Reward{
		Id:                 model.Id,
		TierId:             model.TierId,
		Name:               model.Name,
		Description:        model.Description,
		ExperienceRequired: model.ExperienceRequired,
		Prize:              model.Prize,
		PrizeType:          model.PrizeType,
		Tier:               BuildTierFromModel(model.Tier),
	}
}

func BuildRewardsFromModels(models []models.RewardModel) []Reward {
	var rewards []Reward
	for _, model := range models {
		reward := BuildRewardFromModel(&model)
		rewards = append(rewards, *reward)
	}
	return rewards
}
