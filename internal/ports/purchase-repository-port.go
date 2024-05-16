package ports

import (
	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type PurchaseRepositoryPort interface {
	FindById(id string, where myTypes.Where) (*models.PurchaseModel, error)


}