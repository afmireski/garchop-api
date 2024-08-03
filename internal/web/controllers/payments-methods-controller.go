package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/afmireski/garchop-api/internal/services"
)

type PaymentsMethodsController struct {
	service *services.PaymentsMethodsService
}

func NewPaymentsMethodsController(service *services.PaymentsMethodsService) *PaymentsMethodsController {
	return &PaymentsMethodsController{
		service: service,
	}
}

func (c *PaymentsMethodsController) ListPaymentMethods(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response, err := c.service.ListPaymentMethods()
	if err != nil {
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
