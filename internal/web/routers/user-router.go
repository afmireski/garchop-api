package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupUsersRouter(router *chi.Mux, controller *controllers.UsersController) {
	router.Post("/users/new", controller.NewUser)
	router.Get("/users/{id}", controller.GetUser)
	router.Delete("/users/{id}/del", controller.DeleteClientAccount)
}
