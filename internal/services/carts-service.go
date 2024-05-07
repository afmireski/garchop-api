package services

import (
	"time"

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

func validateAddItemToCartInput(input myTypes.AddItemToCartInput) *customErrors.InternalError {
	if !validators.IsValidUuid(input.UserId) {
		return customErrors.NewInternalError("invalid user_id", 400, []string{"the user_id must be a valid uuid"})
	} else if !validators.IsValidUuid(input.PriceId) {
		return customErrors.NewInternalError("invalid price_id", 400, []string{"the price_id must be a valid uuid"})
	} else if !validators.IsValidUuid(input.PokemonId) {
		return customErrors.NewInternalError("invalid pokemon_id", 400, []string{"the pokemon_id must be a valid uuid"})
	} else if !validators.IsGreaterThanInt[uint](input.Quantity, 0) {
		return customErrors.NewInternalError("invalid quantity", 400, []string{"the quantity must be greater than 0"})
	}

	return nil;
}

func (s *CartsService) AddItemToCart(userId string, input myTypes.AddItemToCartInput) *customErrors.InternalError {

	if inputErr := validateAddItemToCartInput(input); inputErr != nil {
		return inputErr
	}

	cart, findCartErr := s.cartsRepository.FindLastCart(userId); if findCartErr != nil {
		return customErrors.NewInternalError("failed on get the cart", 500, []string{findCartErr.Error()})
	} else if cart == nil {
		newCartInput := myTypes.CreateCartInput{
			UserId: userId,
			IsActive: true,
			ExpiresIn: time.Now().Add(time.Hour * 1),
		}
		newCart, createCartErr := s.cartsRepository.Create(newCartInput); if createCartErr != nil {
			return customErrors.NewInternalError("failed on get the cart", 500, []string{findCartErr.Error()})
		}

		cart = newCart
	}

	

	return nil	
}
