package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/afmireski/garchop-api/internal/web/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/nedpals/supabase-go"
)

func SetupUsersRouter(router *chi.Mux, controller *controllers.UsersController, supabaseClient *supabase.Client) {
	router.Post("/users/new", controller.NewClient)
	router.With(middlewares.SupabaseAuthMiddleware(supabaseClient)).Patch("/users/{id}/update", controller.UpdateClient)
	router.With(middlewares.SupabaseAuthMiddleware(supabaseClient)).Get("/users/{id}", controller.GetUserById)
	router.With(middlewares.SupabaseAuthMiddleware(supabaseClient)).Get("/users/admin", controller.GetAdmins)
	router.With(middlewares.SupabaseAuthMiddleware(supabaseClient)).Delete("/users/{id}/del", controller.DeleteClientAccount)
	router.With(middlewares.SupabaseAuthMiddleware(supabaseClient)).Post("/admin/new", controller.NewAdministrator)
  router.With(middlewares.SupabaseAuthMiddleware(supabaseClient)).Get("/users/{id}/stats", controller.GetUserStatsById)
}
