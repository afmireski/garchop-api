package ports

import (
	"github.com/afmireski/garchop-api/internal/entities"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type RewardsRepositoryPort interface {
	Create(input myTypes.Any) (string, error)	
	FindAll(where myTypes.Where) ([]entities.Reward, error)
	Delete(id string, where myTypes.Where) error
}