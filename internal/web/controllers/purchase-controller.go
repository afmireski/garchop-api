package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/afmireski/garchop-api/internal/services"
	"github.com/go-chi/chi/v5"

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

func (c *PurchaseController) GetPurchasesByUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user_id")

	purchases, err := c.service.GetPurchasesByUser(userId)
	if err != nil {
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(purchases)
}
