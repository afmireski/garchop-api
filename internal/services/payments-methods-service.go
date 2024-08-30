package services

import (
	"github.com/afmireski/garchop-api/internal/entities"
	"github.com/afmireski/garchop-api/internal/ports"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type PaymentsMethodsService struct {
	repository ports.PaymentMethodsRepositoryPort
}

func NewPaymentsMethodsService(repository ports.PaymentMethodsRepositoryPort) *PaymentsMethodsService {
	return &PaymentsMethodsService{
		repository: repository,
	}
}

func (s *PaymentsMethodsService) ListPaymentMethods() ([]entities.PaymentMethod, *customErrors.InternalError) {

	result, err := s.repository.FindAll(myTypes.Where{})

	if err != nil {
		return nil, customErrors.NewInternalError("a failure occurred when try to list the payment methods", 500, []string{err.Error()})
	}

	return entities.BuildManyPaymentMethodsFromModel(result), nil
}
