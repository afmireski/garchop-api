package ports

import (
	"github.com/afmireski/garchop-api/internal/models"

	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type UserRewardsRepositoryPort interface {
	FindAll(where myTypes.Where) ([]models.UserRewardModel, error)

	Create(input UserRewardInput) (models.UserRewardModel, error)

	FindById(input UserRewardInput, where myTypes.Where) (models.UserRewardModel, error)
}

type UserRewardInput struct {
	UserId string
	RewardId string
}