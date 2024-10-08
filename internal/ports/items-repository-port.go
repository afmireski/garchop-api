package ports

import (
	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type ItemsRepositoryPort interface {
	Create(input myTypes.CreateItemInput) (*models.ItemModel, error)
	FindById(id string, where myTypes.Where) (*models.ItemModel, error)
	FindAll(where myTypes.Where) ([]models.ItemModel, error)
	Update(id string, input myTypes.AnyMap, where myTypes.Where) (*models.ItemModel, error)
	UpdateMany(input myTypes.AnyMap, where myTypes.Where) ([]models.ItemModel, error)
	Delete(id string, where myTypes.Where) error
}
