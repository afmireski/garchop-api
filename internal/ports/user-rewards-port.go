package ports

import (
	"github.com/afmireski/garchop-api/internal/models"

	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type UserRewardsRepositoryPort interface {
	FindAll(where myTypes.Where) ([]models.UserRewardModel, error)

	Create(input myTypes.UserRewardInput) (*models.UserRewardModel, error)

	FindById(input myTypes.UserRewardInput, where myTypes.Where) (models.UserRewardModel, error)
}
