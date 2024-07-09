package controllers

import (
	"encoding/json"
	"net/http"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	"github.com/afmireski/garchop-api/internal/services"
	myTypes "github.com/afmireski/garchop-api/internal/types"
	"github.com/go-chi/chi/v5"
)

type RewardsController struct {
	service *services.RewardsService
}

func NewRewardsController(service *services.RewardsService) *RewardsController {
	return &RewardsController{
		service: service,
	}
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

func (c *RewardsController) ClaimReward(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rewardId := chi.URLParam(r, "reward_id")
	userId := r.Header.Get("UserId")

	input := myTypes.UserRewardInput{
		RewardId: rewardId,
		UserId: userId,
	}

	bodyErr := json.NewDecoder(r.Body).Decode(&input); if bodyErr != nil {
		err := customErrors.NewInternalError("fail on deserialize request body", 400, []string{})
		w.WriteHeader(err.HttpCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	serviceErr := c.service.ClaimReward(input)

	if serviceErr != nil {
		w.WriteHeader(serviceErr.HttpCode)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
