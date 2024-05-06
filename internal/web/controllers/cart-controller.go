package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/afmireski/garchop-api/internal/services"
	"github.com/go-chi/chi/v5"
)

type CartController struct {
	service *services.CartsService
}

func NewCartController(service *services.CartsService) *CartController {
	return &CartController{
		service: service,
	}
}

func (c *CartController) GetCurrentUserCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userIdParam := chi.URLParam(r, "user_id")

	response, err := c.service.GetCurrentUserCart(userIdParam); if err != nil {
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}
		
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
