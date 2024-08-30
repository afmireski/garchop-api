package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/afmireski/garchop-api/internal/web/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/nedpals/supabase-go"
)

func SetupPaymentsMethodsRouter(r *chi.Mux, controller *controllers.PaymentsMethodsController, supabaseClient *supabase.Client) {
	r.With(middlewares.SupabaseAuthMiddleware(supabaseClient)).Get("/payments-methods", controller.ListPaymentMethods)
}
