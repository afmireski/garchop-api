package ports

import (
	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type RewardsRepositoryPort interface {
	FindById(id string, where myTypes.Where) (*models.RewardModel, error)
	Create(input myTypes.CreateRewardInput) (string, error)
	FindAll(where myTypes.Where) ([]models.RewardModel, error)
	Update(id string, data myTypes.AnyMap, where myTypes.Where) (*models.RewardModel, error)
}

