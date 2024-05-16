package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupPurchasesRouter(r chi.Router, controller *controllers.PurchaseController) {
	r.Post("/purchases/finish", controller.FinishPurchase)
}
