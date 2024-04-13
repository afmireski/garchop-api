package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupTiersRouter(router *chi.Mux, controller *controllers.TiersController) {
	router.Get("/tiers", controller.FindAll)
	router.Get("/tiers/{id}", controller.FindById)
}
