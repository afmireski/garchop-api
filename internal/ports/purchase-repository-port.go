package ports

import (
	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type PurchaseRepositoryPort interface {
	FindAll(where myTypes.Where) ([]models.PurchaseModel, error)

	FindById(id string, where myTypes.Where) (*models.PurchaseModel, error)

	Create(input myTypes.CreatePurchaseInput) (string, error)

	Delete(id string) error
}