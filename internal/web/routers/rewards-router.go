package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupRewardsRouter(router *chi.Mux, controller *controllers.RewardsController) {
	router.Get("/rewards", controller.ListAllRewards)
	router.Post("/rewards/{reward_id}/claim", controller.ClaimReward)
}