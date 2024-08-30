package controllers

import (
	"encoding/json"
	"net/http"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	"github.com/afmireski/garchop-api/internal/services"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type UsersController struct {
	service *services.UsersService
}

func NewUsersController(service *services.UsersService) *UsersController {
	return &UsersController{
		service,
	}
}

func (c *UsersController) NewClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input myTypes.NewUserInput

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		err := customErrors.NewInternalError("fail on deserialize request body", 400, []string{})
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	serviceErr := c.service.NewClient(input)

	if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *UsersController) NewAdministrator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input myTypes.NewUserInput

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		err := customErrors.NewInternalError("fail on deserialize request body", 400, []string{err.Error()})
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	serviceErr := c.service.NewAdmin(input)

	if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *UsersController) UpdateClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.Header.Get("User-Id")

	var input myTypes.UpdateUserInput

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		err := customErrors.NewInternalError("fail on deserialize request body", 400, []string{})
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	user, serviceErr := c.service.UpdateClient(idParam, input)

	if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (c *UsersController) GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.Header.Get("User-Id")

	user, serviceErr := c.service.GetUserById(idParam)

	if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (c *UsersController) GetAdmins(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, serviceErr := c.service.GetUsers(myTypes.Where{
		"role": map[string]string{"eq": "admin"},
	})

	if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (c *UsersController) GetUserStatsById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.Header.Get("User-Id")

	user, serviceErr := c.service.GetUserStatsById(idParam)

	if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (c *UsersController) DeleteClientAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.Header.Get("User-Id")
	token := r.Header.Get("Authorization")

	serviceErr := c.service.DeleteClient(idParam, token)

	if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
