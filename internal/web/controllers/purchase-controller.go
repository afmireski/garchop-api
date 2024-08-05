package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/afmireski/garchop-api/internal/services"

	myTypes "github.com/afmireski/garchop-api/internal/types"
	customErrors "github.com/afmireski/garchop-api/internal/errors"
)

type PurchaseController struct {
	service *services.PurchasesService
}

func NewPurchaseController(service *services.PurchasesService) *PurchaseController {
	return &PurchaseController{
		service: service,
	}
}

func (c *PurchaseController) FinishPurchase(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("User-Id")

	var body myTypes.FinishPurchaseBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customErrors.NewInternalError("fail on deserialize request body", 400, []string{err.Error()}))
		return
	}

	input := myTypes.FinishPurchaseInput{
		UserId: userId,
		CartId: body.CartId,
		PaymentMethodId: body.PaymentMethodId,
	}

	serviceErr := c.service.FinishPurchase(input); if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (c *PurchaseController) GetPurchasesByUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("User-Id")

	purchases, err := c.service.GetPurchasesByUser(userId)
	if err != nil {
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	status := http.StatusOK
	if len(purchases) == 0 {
		status = http.StatusNoContent
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(purchases)
}
