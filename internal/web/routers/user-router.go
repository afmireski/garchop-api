package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupUsersRouter(router *chi.Mux, controller *controllers.UsersController) {
	router.Post("/users/new", controller.NewClient)
	router.Patch("/users/{id}/update", controller.UpdateClient)
	router.Get("/users/{id}", controller.GetUserById)
	router.Delete("/users/{id}/del", controller.DeleteClientAccount)

	router.Post("/admin/new", controller.NewAdministrator)
}
