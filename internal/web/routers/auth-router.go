package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupAuthRouter(router *chi.Mux, controller *controllers.AuthController) {
	router.Post("/auth/signin", controller.SignIn)
}