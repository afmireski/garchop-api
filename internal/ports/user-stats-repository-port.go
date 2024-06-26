package ports

import (
	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type UserStatsRepository interface {
	Create(input myTypes.CreateUserStatsInput) (string, error)
	FindById(id string, where myTypes.Where) (*models.UserStatsModel, error)
	Update(id string, input myTypes.AnyMap, where myTypes.Where) (*models.UserStatsModel, error)
	Delete(id string, where myTypes.Where) error
}
