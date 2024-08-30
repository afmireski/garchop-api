package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/afmireski/garchop-api/internal/web/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/nedpals/supabase-go"
)

func SetupCartsRouter(r chi.Router, controller *controllers.CartController, supabaseClient *supabase.Client) {
	r.With(middlewares.SupabaseAuthMiddleware(supabaseClient)).Get("/cart", controller.GetCurrentUserCart)
	r.With(middlewares.SupabaseAuthMiddleware(supabaseClient)).Post("/cart/add-item", controller.AddItemToCart)
}
