package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/afmireski/garchop-api/internal/web/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/nedpals/supabase-go"
)

func SetupItemsRouter(r chi.Router, controller *controllers.ItemController, supabaseClient *supabase.Client) {
	r.With(middlewares.SupabaseAuthMiddleware(supabaseClient), middlewares.UserRoleMiddleware("client")).Delete("/carts/{cart_id}/items/{item_id}", controller.RemoveItemFromCart)
	r.With(middlewares.SupabaseAuthMiddleware(supabaseClient), middlewares.UserRoleMiddleware("client")).Patch("/carts/{cart_id}/items/{item_id}", controller.UpdateItemInCart)
}
