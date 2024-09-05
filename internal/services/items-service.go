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
	stockRepository ports.StockRepositoryPort
}

func NewItemsService(itemsRepository ports.ItemsRepositoryPort, cartsRepository ports.CartsRepositoryPort, stockRepository ports.StockRepositoryPort) *ItemsService {
	return &ItemsService{
		itemsRepository: itemsRepository,
		cartsRepository: cartsRepository,
		stockRepository: stockRepository,
	}
}

func (s *ItemsService) RemoveItemFromCart(input myTypes.RemoveItemFromCartInput) *customErrors.InternalError {
	if !validators.IsValidUuid(input.ItemId) {
		return customErrors.NewInternalError("invalid item_id", 400, []string{"the item_id must be a valid uuid"})
	}

	if !validators.IsValidUuid(input.CartId) {
		return customErrors.NewInternalError("invalid cart_id", 400, []string{"the cart_id must be a valid uuid"})
	}

	itemWhere := myTypes.Where{
		"cart_id": map[string]string{"eq": input.CartId},
	}

	tmpItemData, itemErr := s.itemsRepository.FindById(input.ItemId, itemWhere)

	if itemErr != nil {
		return customErrors.NewInternalError("a failure occurred when try to remove the item", 500, []string{itemErr.Error()})
	} else if tmpItemData == nil {
		return customErrors.NewInternalError("a failure occurred when try to remove the item", 404, []string{"the item was not found"})
	}

	cartData, cartErr := s.cartsRepository.FindById(input.CartId, nil)

	if cartErr != nil {
		return customErrors.NewInternalError("a failure occurred when try to remove the item", 500, []string{cartErr.Error()})
	} else if cartData == nil {
		return customErrors.NewInternalError("a failure occurred when try to remove the item", 404, []string{"the cart was not found"})
	}

	stockData, _ := s.stockRepository.FindById(tmpItemData.PokemonId, myTypes.Where{"deleted_at": map[string]string{"is": "null"}})

	err := s.itemsRepository.Delete(input.ItemId, itemWhere)

	if err != nil {
		return customErrors.NewInternalError("failed on delete the item", 500, []string{err.Error()})
	}

	_, err = s.cartsRepository.Update(input.CartId, myTypes.AnyMap{
		"total": cartData.Total - tmpItemData.Total,
	}, nil)

	if err != nil {
		return customErrors.NewInternalError("failed on update the cart", 500, []string{err.Error()})
	}

	_, err = s.stockRepository.Update(tmpItemData.PokemonId, myTypes.AnyMap{"quantity": stockData.Quantity + tmpItemData.Quantity}, myTypes.Where{
		"deleted_at": map[string]string{"is": "null"},
	})

	if err != nil {
		return customErrors.NewInternalError("failed on update the stock", 500, []string{err.Error()})
	}

	return nil

}

func (s *ItemsService) validateUpdateItemInCart(input myTypes.UpdateItemInCartInput) *customErrors.InternalError {
	if !validators.IsValidUuid(input.ItemId) {
		return customErrors.NewInternalError("invalid item_id", 400, []string{"the item_id must be a valid uuid"})
	}

	if !validators.IsValidUuid(input.CartId) {
		return customErrors.NewInternalError("invalid cart_id", 400, []string{"the cart_id must be a valid uuid"})
	}

	if !validators.IsGreaterThanInt(input.Quantity, 1) {
		return customErrors.NewInternalError("invalid quantity", 400, []string{"the quantity must be greater than 1"})
	}

	return nil
}

func (s *ItemsService) UpdateItemInCart(input myTypes.UpdateItemInCartInput) *customErrors.InternalError {

	inputErr := s.validateUpdateItemInCart(input); if inputErr != nil {
		return inputErr
	}

	findWhere := myTypes.Where{
		"cart_id": map[string]string{"eq": input.CartId},
		"deleted_at": map[string]string{"is": "null"},
		"purchase_id": map[string]string{"is": "null"},
	}
	item, findErr := s.itemsRepository.FindById(input.ItemId, findWhere); if findErr != nil {
		return customErrors.NewInternalError("a failure occurred when try to update the item", 500, []string{findErr.Error()})
	}

	// Se a quantidade do item não mudou, não faz nada.
	if input.Quantity == item.Quantity {
		return nil
	}

	remainStock := item.Pokemon.Stock.Quantity
	// Se a quantidade do item aumentou, então precisa verificar o estoque
	if input.Quantity > item.Quantity  {
		deltaQuantity := input.Quantity - item.Quantity // Descobre quanto aumentou
		remainStock -= deltaQuantity // Calcula quanto de estoque vai sobrar

		if remainStock < 0 {
			return customErrors.NewInternalError("the quantity is greater than the stock", 400, []string{"the quantity is greater than the stock"})
		}
	} else {
		// Se a quantidade do item diminuiu, então precisa repor o estoque

		deltaQuantity := item.Quantity - input.Quantity // Descobre quanto diminuiu
		remainStock += deltaQuantity // Calcula quanto de estoque vai sobrar
	}

	_, stockErr := s.stockRepository.Update(item.PokemonId, myTypes.AnyMap{"quantity": remainStock}, myTypes.Where{
		"deleted_at": map[string]string{"is": "null"},
	}); if stockErr != nil {
		return customErrors.NewInternalError("a failure occurred when try to update the stock", 500, []string{stockErr.Error()})
	}


	data := myTypes.AnyMap{
		"quantity": input.Quantity,
		"total": item.Price * input.Quantity,
	}
	where := myTypes.Where{
		"deleted_at": map[string]string{"is": "null"},
	}
	_, err := s.itemsRepository.Update(input.ItemId, data, where); if err != nil {
		return customErrors.NewInternalError("a failure occurred when try to update the item", 500, []string{err.Error()})
	}


	return nil
}
