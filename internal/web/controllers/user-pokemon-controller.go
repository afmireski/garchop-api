package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/afmireski/garchop-api/internal/services"
	myTypes "github.com/afmireski/garchop-api/internal/types"
	"github.com/go-chi/chi/v5"
)

type UserPokemonController struct {
	service services.UserPokemonsService
}

func NewUserPokemonController(service services.UserPokemonsService) *UserPokemonController {
	return &UserPokemonController{
		service: service,
	}
}

func (c *UserPokemonController) GetUserPokedex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId := r.Header.Get("User-Id")

	input := myTypes.GetPokedexInput{
		UserId: userId,
	}

	response, serviceErr := c.service.GetUserPokedex(input); if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *UserPokemonController) GetUserPokemonDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId := r.Header.Get("User-Id")
	pokemonId := chi.URLParam(r, "pokemonId")

	input := myTypes.GetUserPokemonInput{
		UserId: userId,
		PokemonId: pokemonId,
	}

	response, serviceErr := c.service.GetUserPokemon(input); if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

