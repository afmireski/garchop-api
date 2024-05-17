package services

import (
	"time"

	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/validators"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type PurchasesService struct {
	repository     ports.PurchaseRepositoryPort
	cartRepository ports.CartsRepositoryPort
}

func NewPurchasesService(repository ports.PurchaseRepositoryPort, cartRepository ports.CartsRepositoryPort) *PurchasesService {
	return &PurchasesService{
		repository:     repository,
		cartRepository: cartRepository,
	}
}

func validateFinishPurchaseInput(input myTypes.FinishPurchaseInput) *customErrors.InternalError {

	if !validators.IsValidUuid(input.UserId) {
		return customErrors.NewInternalError("invalid user_id", 400, []string{"the user_id must be a valid uuid"})
	} else if !validators.IsValidUuid(input.CartId) {
		return customErrors.NewInternalError("invalid cart_id", 400, []string{"the cart_id must be a valid uuid"})
	} else if !validators.IsValidUuid(*input.PaymentMethodId) {
		return customErrors.NewInternalError("invalid payment_method_id", 400, []string{"the payment_method_id must be a valid uuid"})
	}

	return nil
}

func (s *PurchasesService) FinishPurchase(input myTypes.FinishPurchaseInput) *customErrors.InternalError {
	err := validateFinishPurchaseInput(input)
	if err != nil {
		return err
	}

	cart, findCartErr := s.cartRepository.FindById(input.CartId, myTypes.Where{
		"user_id":    map[string]string{"eq": input.UserId},
		"is_active":  map[string]string{"is": "true"},
		"expires_in": map[string]string{"gt": "now()"},
		"deleted_at": map[string]string{"is": "null"},
	})
	if findCartErr != nil {
		return customErrors.NewInternalError("failed on get the cart", 500, []string{err.Error()})
	} else if cart == nil {
		return customErrors.NewInternalError("cart not found", 404, []string{})
	}

	data := myTypes.CreatePurchaseInput{
		CartId:           input.CartId,
		UserId:           input.UserId,
		PaymentMethodId:  input.PaymentMethodId,
		Total:            cart.Total,
		PaymentLimitDate: time.Now().Add(time.Minute * 30),
	}

	_, finishErr := s.repository.Create(data); if finishErr != nil {
		return customErrors.NewInternalError("failed on finish the purchase", 500, []string{finishErr.Error()})
	}

	deleteCartErr := s.cartRepository.Delete(input.CartId); if deleteCartErr != nil {
		return customErrors.NewInternalError("failed on delete the cart", 500, []string{deleteCartErr.Error()})
	}

	return nil
}
