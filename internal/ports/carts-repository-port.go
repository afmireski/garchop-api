package ports

import (
	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type CartsRepositoryPort interface {
	Create(input myTypes.Any) (string, error)

	FindById(id string, where myTypes.Where, order myTypes.Order) ([]myTypes.AnyMap, error)

	FindLastCart(user_id string) (*models.CartModel, error)

	Update(id string, input myTypes.AnyMap, where myTypes.Where) (*myTypes.AnyMap, error)

	Delete(id string) error
}
