package ports

import (
	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type CartsRepositoryPort interface {
	Create(input myTypes.CreateCartInput) (*models.CartModel, error)

	FindById(id string, where myTypes.Where) (*models.CartModel, error)

	FindLastCart(user_id string) (*models.CartModel, error)

	Update(id string, input myTypes.AnyMap, where myTypes.Where) (*models.CartModel, error)

	Delete(id string) error
}
