package entities

import "github.com/afmireski/garchop-api/internal/models"

type UserStats struct {
	UserId     string `json:"user_id"`
	TierId     uint   `json:"tier_id"`
	Experience uint   `json:"experience"`
	Tier       *Tier  `json:"tiers",omitempty`
}

func BuildUserStatsFromModel(model *models.UserStatsModel) *UserStats {
	if model == nil {
		return nil
	}

	return &UserStats{
		UserId:     model.UserId,
		TierId:     model.TierId,
		Experience: model.Experience,
		Tier:       BuildTierFromModel(model.Tier),
	}
}
