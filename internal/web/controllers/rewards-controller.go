package controllers

import (
	"encoding/json"
	"net/http"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	"github.com/afmireski/garchop-api/internal/services"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type RewardsController struct {
	service *services.RewardsService
}

func NewRewardsController(service *services.RewardsService) *RewardsController {
	return &RewardsController{
		service: service,
	}
}

func (c *RewardsController) NewReward(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input myTypes.NewRewardInput

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		err := customErrors.NewInternalError("fail on deserialize request body", 400, []string{})
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	serviceErr := c.service.NewReward(input)

	if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *RewardsController) ListAllRewards(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response, err := c.service.ListAllRewards()

	if err != nil {
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	if len(response) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}