package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/afmireski/garchop-api/internal/web/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/nedpals/supabase-go"
)

func SetupRewardsRouter(
	router *chi.Mux, 
	controller *controllers.RewardsController, 
	supabaseClient *supabase.Client,
	) {
	router.Get("/rewards", controller.ListAllRewards)
	router.With(
		middlewares.SupabaseAuthMiddleware(supabaseClient), 
		middlewares.UserRoleMiddleware("client")).Post("/rewards/{reward_id}/claim", controller.ClaimReward)
	router.With(
		middlewares.SupabaseAuthMiddleware(supabaseClient), 
		middlewares.UserRoleMiddleware("admin")).Post("/rewards/new", controller.NewReward)
}