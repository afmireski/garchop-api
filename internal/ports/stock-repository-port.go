package ports

import (
	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type StockRepositoryPort interface {
	FindById(id string, where myTypes.Where) (*models.StockModel, error)
	Update(id string, input myTypes.AnyMap, where myTypes.Where) (*models.StockModel, error)
}