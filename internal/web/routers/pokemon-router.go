package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupPokemonRouter(r chi.Router, controller *controllers.PokemonController) {
	r.Post("/pokemons/new", controller.RegistryNewPokemon)
}