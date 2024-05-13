package services

import (
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/validators"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type ItemsService struct {
	itemsRepository ports.ItemsRepositoryPort
	cartsRepository ports.CartsRepositoryPort
}

func NewItemsService(itemsRepository ports.ItemsRepositoryPort, cartsRepository ports.CartsRepositoryPort) *ItemsService {
	return &ItemsService{
		itemsRepository: itemsRepository,
		cartsRepository: cartsRepository,
	}
}

func (s *ItemsService) RemoveItemFromCart(input myTypes.RemoveItemFromCartInput) *customErrors.InternalError {
	if !validators.IsValidUuid(input.ItemId) {
		return customErrors.NewInternalError("invalid user_id", 400, []string{"the item_id must be a valid uuid"})
	}

	if !validators.IsValidUuid(input.CartId) {
		return customErrors.NewInternalError("invalid item_id", 400, []string{"the cart_id must be a valid uuid"})
	}

	itemWhere := myTypes.Where{
		"cart_id": map[string]string{"eq": input.CartId},
	}

	tmpItem, err := s.itemsRepository.FindById(input.ItemId, itemWhere)

	if err != nil {
		return customErrors.NewInternalError("failed on get the item", 500, []string{err.Error()})
	} else if tmpItem == nil {
		return customErrors.NewInternalError("item not found", 404, []string{})
	}

	cartData, err := s.cartsRepository.FindById(input.CartId, nil)

	if err != nil {
		return customErrors.NewInternalError("failed on get the cart", 500, []string{err.Error()})
	} else if cartData == nil {
		return customErrors.NewInternalError("cart not found", 404, []string{})
	}

	err = s.itemsRepository.Delete(input.ItemId)

	if err != nil {
		return customErrors.NewInternalError("failed on delete the item", 500, []string{err.Error()})
	}

	_, err = s.cartsRepository.Update(input.CartId, myTypes.AnyMap{
		"total": cartData.Total - tmpItem.Total,
	}, nil)

	if err != nil {
		return customErrors.NewInternalError("failed on update the cart", 500, []string{err.Error()})
	}

	return nil

}
