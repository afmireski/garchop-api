package entities

import "github.com/afmireski/garchop-api/internal/models"

type PaymentMethod struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func BuildPaymentMethodFromModel(model models.PaymentMethodModel) PaymentMethod {
	return PaymentMethod{
		Id:   model.Id,
		Name: model.Name,
	}
}

func BuildManyPaymentMethodsFromModel(models []models.PaymentMethodModel) []PaymentMethod {
	var paymentMethods []PaymentMethod
	for _, model := range models {
		paymentMethods = append(paymentMethods, BuildPaymentMethodFromModel(model))
	}
	return paymentMethods
}
