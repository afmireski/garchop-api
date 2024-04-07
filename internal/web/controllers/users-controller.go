package controllers

import (
	"encoding/json"
	"net/http"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	"github.com/afmireski/garchop-api/internal/services"
	myTypes "github.com/afmireski/garchop-api/internal/types"
	"github.com/go-chi/chi/v5"
)

type UsersController struct {
	service *services.UsersService
}

func NewUsersController(service *services.UsersService) *UsersController {
	return &UsersController{
		service,
	}
}

func (c *UsersController) NewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input myTypes.NewUserInput

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		err := customErrors.NewInternalError("fail on deserialize request body", 400, []string{})
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	serviceErr := c.service.NewUser(input)

	if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *UsersController) UpdateClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")

	var input myTypes.UpdateUserInput

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		err := customErrors.NewInternalError("fail on deserialize request body", 400, []string{})
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	user := c.service.UpdateClient(idParam, input)

	if user != nil {
		w.WriteHeader(user.HttpCode)
		json.NewEncoder(w).Encode(user)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *UsersController) GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")

	user, serviceErr := c.service.GetUserById(idParam)

	if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (c *UsersController) DeleteClientAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")

	serviceErr := c.service.DeleteClient(idParam)

	if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
