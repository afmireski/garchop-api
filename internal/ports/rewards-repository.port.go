package ports

import (
	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type RewardsRepositoryPort interface {
	FindAll(where myTypes.Where, pagination myTypes.Pagination) ([]models.RewardModel, error)
	FindById(id string, where myTypes.Where) (*models.RewardModel, error)
	Create(input myTypes.CreateRewardInput) (string, error)
	Update(id string, data myTypes.AnyMap, where myTypes.Where) (*models.RewardModel, error)
}

