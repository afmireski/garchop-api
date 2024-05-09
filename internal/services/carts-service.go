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
	priceRepository ports.PriceRepositoryPort
	stockRepository ports.StockRepositoryPort
}

func NewCartsService(cartsRepository ports.CartsRepositoryPort, itemsRepository ports.ItemsRepositoryPort, priceRepository ports.PriceRepositoryPort, stockRepository ports.StockRepositoryPort) *CartsService {
	return &CartsService{
		cartsRepository: cartsRepository,
		itemsRepository: itemsRepository,
		priceRepository: priceRepository,
		stockRepository: stockRepository,
	}
}

func (s *CartsService) GetCurrentUserCart(user_id string) (*entities.Cart, *customErrors.InternalError) {
	if !validators.IsValidUuid(user_id) {
		return nil, customErrors.NewInternalError("invalid user_id", 400, []string{"the user_id must be a valid uuid"})
	}

	cartsRepositoryData, crErr := s.cartsRepository.FindLastCart(user_id)
	if crErr != nil {
		return nil, customErrors.NewInternalError("failed on get the cart", 500, []string{crErr.Error()})
	} else if cartsRepositoryData == nil {
		return nil, customErrors.NewInternalError("cart not found", 404, []string{})
	}

	irWhere := myTypes.Where{
		"cart_id":    map[string]string{"eq": cartsRepositoryData.Id},
		"deleted_at": map[string]string{"is": "null"},
	}
	itensRepositoryData, irErr := s.itemsRepository.FindAll(irWhere)
	if irErr != nil {
		return nil, customErrors.NewInternalError("failed on get the cart", 500, []string{irErr.Error()})
	}

	cartsRepositoryData.Items = itensRepositoryData

	return entities.BuildCartFromModel(*cartsRepositoryData), nil
}

func validateAddItemToCartInput(input myTypes.AddItemToCartInput) *customErrors.InternalError {
	if !validators.IsValidUuid(input.UserId) {
		return customErrors.NewInternalError("invalid user_id", 400, []string{"the user_id must be a valid uuid"})
	} else if !validators.IsValidUuid(input.PokemonId) {
		return customErrors.NewInternalError("invalid pokemon_id", 400, []string{"the pokemon_id must be a valid uuid"})
	} else if !validators.IsGreaterThanInt[uint](input.Quantity, 0) {
		return customErrors.NewInternalError("invalid quantity", 400, []string{"the quantity must be greater than 0"})
	}

	return nil
}

func (s *CartsService) AddItemToCart(input myTypes.AddItemToCartInput) *customErrors.InternalError {

	if inputErr := validateAddItemToCartInput(input); inputErr != nil {
		return inputErr
	}

	// Find Pokemon Price
	price, priceErr := s.priceRepository.FindCurrentPrice(input.PokemonId)
	if priceErr != nil {
		return customErrors.NewInternalError("failed on get pokemon infos", 500, []string{priceErr.Error()})
	}

	if price.Pokemons.Stock.Quantity < input.Quantity || price.Pokemons.Stock.Quantity == 0 {
		return customErrors.NewInternalError("this pokemon is out of stock", 400, []string{priceErr.Error()})
	}

	// Get the user cart
	cart, findCartErr := s.cartsRepository.FindLastCart(input.UserId)
	if findCartErr != nil {
		return customErrors.NewInternalError("failed on get the cart", 500, []string{findCartErr.Error()})
	} else if cart == nil {
		newCartInput := myTypes.CreateCartInput{
			UserId:    input.UserId,
			IsActive:  true,
			ExpiresIn: time.Now().Add(time.Hour * 1),
		}
		newCart, createCartErr := s.cartsRepository.Create(newCartInput)
		if createCartErr != nil {
			return customErrors.NewInternalError("failed on get the cart", 500, []string{findCartErr.Error()})
		}

		cart = newCart
	}

	itemInput := myTypes.CreateItemInput{
		CartId:    cart.Id,
		PokemonId: input.PokemonId,
		Price:     price.Value,
		Quantity:  input.Quantity,
		Total:     price.Value * input.Quantity,
	}

	s.itemsRepository.Create(itemInput)

	cartWhere := myTypes.Where{
		"id":         map[string]string{"eq": cart.Id},
		"deleted_at": map[string]string{"is": "null"},
	}

	cartData := myTypes.AnyMap{
		"total": cart.Total + price.Value*input.Quantity,
	}

	_, updateCartErr := s.cartsRepository.Update(cart.Id, cartData, cartWhere)
	if updateCartErr != nil {
		return customErrors.NewInternalError("failed on update cart total", 500, []string{updateCartErr.Error()})
	}

	s.stockRepository.Update(input.PokemonId, myTypes.AnyMap{"quantity": price.Pokemons.Stock.Quantity - input.Quantity}, myTypes.Where{
		"deleted_at": map[string]string{"is": "null"},		
	})

	return nil
}
