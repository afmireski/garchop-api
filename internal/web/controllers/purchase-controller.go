package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/afmireski/garchop-api/internal/services"

	myTypes "github.com/afmireski/garchop-api/internal/types"
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

	var input myTypes.FinishPurchaseInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	serviceErr := c.service.FinishPurchase(input); if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
