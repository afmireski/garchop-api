package ports

import (
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type CartsRepositoryPort interface {
	Create(input myTypes.Any) (string, error)

	FindById(id string, where myTypes.Where, order myTypes.Order) ([]myTypes.AnyMap, error)

	FindAll(where myTypes.Where, order myTypes.Order) ([]myTypes.AnyMap, error)

	Update(id string, input myTypes.AnyMap, where myTypes.Where) (*myTypes.AnyMap, error)

	Delete(id string) error
}
