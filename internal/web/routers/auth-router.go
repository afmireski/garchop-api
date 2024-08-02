package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/afmireski/garchop-api/internal/web/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/nedpals/supabase-go"
)

func SetupAuthRouter(router *chi.Mux, controller *controllers.AuthController, supabaseClient *supabase.Client) {
	router.Post("/sign-in", controller.SignIn)
	router.With(middlewares.SupabaseAuthMiddleware(supabaseClient)).Post("/sign-out", controller.SignOut)
}