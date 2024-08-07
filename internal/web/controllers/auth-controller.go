package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/afmireski/garchop-api/internal/services"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{
		service,
	}
}

func (c *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var credentials myTypes.LoginInput

	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		err := customErrors.NewInternalError("fail on deserialize request body", 400, []string{})
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	response, serviceErr := c.service.Login(credentials)

	if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *AuthController) SignOut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	authorizationToken := r.Header.Get("Authorization")

	err := c.service.Logout(authorizationToken); if err != nil {
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}