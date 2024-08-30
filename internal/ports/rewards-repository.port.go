package ports

import (
	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type RewardsRepositoryPort interface {
	Create(input myTypes.Any) (string, error)	
	FindAll(where myTypes.Where, pagination myTypes.Pagination) ([]models.RewardModel, error)
	Delete(id string, where myTypes.Where) error
}