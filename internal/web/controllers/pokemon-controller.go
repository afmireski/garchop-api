package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/afmireski/garchop-api/internal/helpers"
	"github.com/afmireski/garchop-api/internal/services"
	"github.com/go-chi/chi/v5"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type PokemonController struct {
	service *services.PokemonService
}

func NewPokemonController(service *services.PokemonService) *PokemonController {
	return &PokemonController{
		service: service,
	}
}

func (c *PokemonController) RegistryNewPokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input myTypes.NewPokemonInput
	err := json.NewDecoder(r.Body).Decode(&input); if err != nil {
		err := customErrors.NewInternalError("fail on deserialize request body", 400, []string{})
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	serviceErr := c.service.NewPokemon(input); if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *PokemonController) GetAllPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var filter myTypes.Where
	filterErr := helpers.ParseQueryParam(r.URL.Query(), "filter", &filter)

	if filterErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customErrors.NewInternalError("fail on deserialize query parameters", 400, []string{}))
		return
	}

	response, err := c.service.GetAvailablePokemons(filter)
	if err != nil {
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *PokemonController) GetPokemonById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")

	response, err := c.service.GetPokemonById(idParam); if err != nil {
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}
		
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *PokemonController) DeletePokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")

	err := c.service.DeletePokemon(idParam); if err != nil {
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *PokemonController) UpdatePokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")

	var input myTypes.UpdatePokemonInput
	err := json.NewDecoder(r.Body).Decode(&input); if err != nil {
		err := customErrors.NewInternalError("fail on deserialize request body", 400, []string{})
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	response, serviceErr := c.service.UpdatePokemon(idParam, input); if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
