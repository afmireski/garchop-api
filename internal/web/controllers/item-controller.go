package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/afmireski/garchop-api/internal/services"
	myTypes "github.com/afmireski/garchop-api/internal/types"
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
