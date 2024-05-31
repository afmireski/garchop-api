package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/afmireski/garchop-api/internal/web/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/nedpals/supabase-go"
)

func SetupPokemonRouter(r chi.Router, controller *controllers.PokemonController, supabaseClient *supabase.Client) {
	r.With(middlewares.SupabaseAuthMiddleware(supabaseClient)).Post("/pokemon/new", controller.RegistryNewPokemon)
	r.Get("/pokemon/{id}", controller.GetPokemonById)
	r.Get("/pokemon", controller.GetAllPokemons)
	r.With(middlewares.SupabaseAuthMiddleware(supabaseClient)).Patch("/pokemon/{id}/update", controller.UpdatePokemon)
	r.With(middlewares.SupabaseAuthMiddleware(supabaseClient)).Delete("/pokemon/{id}/del", controller.DeletePokemon)
}