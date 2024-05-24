package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupItemsRouter(r chi.Router, controller *controllers.ItemController) {
	r.Delete("/carts/{cart_id}/items/{item_id}", controller.RemoveItemFromCart)
}
