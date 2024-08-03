package ports

import (
	"github.com/afmireski/garchop-api/internal/models"

	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type PaymentMethodsRepositoryPort interface {
	Create(name string) (*models.PaymentMethodModel, error)
	FindAll(where myTypes.Where) ([]models.PaymentMethodModel, error)
}
