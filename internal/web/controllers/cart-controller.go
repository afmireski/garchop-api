package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/afmireski/garchop-api/internal/services"
	myTypes "github.com/afmireski/garchop-api/internal/types"
	customErrors "github.com/afmireski/garchop-api/internal/errors"
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

	userId := r.Header.Get("User-Id")

	response, err := c.service.GetCurrentUserCart(userId); if err != nil {
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}
		
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *CartController) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body myTypes.AddItemToCartBody
	bodyErr := json.NewDecoder(r.Body).Decode(&body); if bodyErr != nil {
		err := customErrors.NewInternalError("fail on deserialize request body", 400, []string{})
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	userId := r.Header.Get("User-Id")

	input := myTypes.AddItemToCartInput{
		UserId: userId,
		PokemonId: body.PokemonId,
		Quantity: body.Quantity,
	}

	err := c.service.AddItemToCart(input); if err != nil {
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
