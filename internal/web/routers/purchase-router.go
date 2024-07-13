package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/afmireski/garchop-api/internal/web/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/nedpals/supabase-go"
)

func SetupPurchasesRouter(r chi.Router, controller *controllers.PurchaseController, supabaseClient *supabase.Client) {
	r.With(middlewares.SupabaseAuthMiddleware(supabaseClient)).Post("/purchases/finish", controller.FinishPurchase)
	r.Get("/users/{user_id}/purchases", controller.GetPurchasesByUser)
}
