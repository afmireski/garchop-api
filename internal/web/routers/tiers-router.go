package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/afmireski/garchop-api/internal/web/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/nedpals/supabase-go"
)

func SetupTiersRouter(router *chi.Mux, controller *controllers.TiersController, supabaseClient *supabase.Client) {
	router.With(middlewares.SupabaseAuthMiddleware(supabaseClient)).Get("/tiers", controller.FindAll)
	router.Get("/tiers/{id}", controller.FindById)
}
