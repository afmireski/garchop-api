package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/afmireski/garchop-api/internal/services"
	"github.com/go-chi/chi/v5"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
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
	id, err := strconv.Atoi(idParam)

	if err != nil {
		err := customErrors.NewInternalError("invalid parameter", 400, []string{"failure on try serialize id parameter"})
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	response, serviceErr := c.service.FindById(id); if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}
		
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
