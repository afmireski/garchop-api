package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/afmireski/garchop-api/internal/services"
	"github.com/go-chi/chi/v5"
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

func (c *TiersController) FindById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")

	response, err := c.service.FindById(idParam); if err != nil {
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}
		
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
