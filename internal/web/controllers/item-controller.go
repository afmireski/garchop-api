package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/afmireski/garchop-api/internal/services"
	myTypes "github.com/afmireski/garchop-api/internal/types"
	customErrors "github.com/afmireski/garchop-api/internal/errors"

	"github.com/go-chi/chi/v5"
)

type ItemController struct {
	service *services.ItemsService
}

func NewItemController(service *services.ItemsService) *ItemController {
	return &ItemController{
		service: service,
	}
}

func (c *ItemController) RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	itemId := chi.URLParam(r, "item_id")
	cartId := chi.URLParam(r, "cart_id")

	input := myTypes.RemoveItemFromCartInput{
		ItemId: itemId,
		CartId: cartId,
	}

	err := c.service.RemoveItemFromCart(input)
	if err != nil {
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *ItemController) UpdateItemInCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	itemId := chi.URLParam(r, "item_id")
	cartId := chi.URLParam(r, "cart_id")

	var body myTypes.UpdateItemInCartBody
	bodyErr := json.NewDecoder(r.Body).Decode(&body); if bodyErr != nil {
		err := customErrors.NewInternalError("fail on deserialize request body", 400, []string{bodyErr.Error()})
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	input := myTypes.UpdateItemInCartInput{
		ItemId: itemId,
		CartId: cartId,
		Quantity: body.Quantity,
	}

	err := c.service.UpdateItemInCart(input)
	if err != nil {
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

