package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupCartsRouter(r chi.Router, controller *controllers.CartController) {
	r.Get("/users/{user_id}/cart", controller.GetCurrentUserCart)
	r.Post("/users/{user_id}/cart/add-item", controller.AddItemToCart)
}