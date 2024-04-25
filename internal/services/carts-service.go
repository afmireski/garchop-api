package services

import (
	"github.com/afmireski/garchop-api/internal/entities"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/validators"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type CartsService struct {
	cartsRepository ports.CartsRepositoryPort
	itemsRepository ports.ItemsRepositoryPort
}

func NewCartsService(cartsRepository ports.CartsRepositoryPort, itemsRepository ports.ItemsRepositoryPort) *CartsService {
	return &CartsService{
		cartsRepository: cartsRepository,
		itemsRepository: itemsRepository,
	}
}

func (s *CartsService) GetCurrentUserCart(user_id string) (*entities.Cart, *customErrors.InternalError) {
	if !validators.IsValidUuid(user_id) {
		return nil, customErrors.NewInternalError("invalid user_id", 400, []string{"the user_id must be a valid uuid"})
	}

	cartsRepositoryData, crErr := s.cartsRepository.FindLastCart(user_id); if crErr != nil {
		return nil, customErrors.NewInternalError("failed on get the cart", 500, []string{crErr.Error()})
	} else if cartsRepositoryData == nil {
		return nil, customErrors.NewInternalError("cart not found", 404, []string{})
	}

	irWhere := myTypes.Where{
		"cart_id": map[string]string{"eq": cartsRepositoryData.Id},
		"deleted_at": map[string]string{"is": "null"},
	}
	itensRepositoryData, irErr := s.itemsRepository.FindAll(irWhere); if irErr != nil {
		return nil, customErrors.NewInternalError("failed on get the cart", 500, []string{irErr.Error()})
	}

	cartsRepositoryData.Items = itensRepositoryData

	return entities.BuildCartFromModel(*cartsRepositoryData), nil
}
