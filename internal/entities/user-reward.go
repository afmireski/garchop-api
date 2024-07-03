package entities

import (
	"time"

	"github.com/afmireski/garchop-api/internal/models"
)

type UserReward struct {
	UserId    string     `json:"user_id"`
	RewardId  string     `json:"reward_id"`
	ClaimedAt time.Time `json:"claimed_at"`
	Reward    *Reward    `json:"reward"`
}

func BuildUserRewardFromModel(model *models.UserRewardModel) *UserReward {
	return &UserReward{
		UserId:    model.UserId,
		RewardId:  model.RewardId,
		ClaimedAt: model.ClaimedAt,
		Reward:    BuildRewardFromModel(model.RewardModel),
	}
}
