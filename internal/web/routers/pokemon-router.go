package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupPokemonRouter(r chi.Router, controller *controllers.PokemonController) {
	r.Post("/pokemon/new", controller.RegistryNewPokemon)
	r.Get("/pokemon/{id}", controller.GetPokemonById)
	r.Get("/pokemon", controller.GetAllPokemons)
	r.Delete("/pokemon/{id}/del", controller.DeletePokemon)
}