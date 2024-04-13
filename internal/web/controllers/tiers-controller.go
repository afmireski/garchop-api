package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/afmireski/garchop-api/internal/services"
)

type TiersController struct {
	service *services.TiersService
}

func NewTiersController(service *services.TiersService) *TiersController {
	return &TiersController{
		service: service,
	}
}

func (c *TiersController) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response, err := c.service.FindAll(); if err != nil {
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
