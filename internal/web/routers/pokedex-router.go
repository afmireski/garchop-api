package routers

import (
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/afmireski/garchop-api/internal/web/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/nedpals/supabase-go"
)

func SetupPokedexRouter(
	r chi.Router,
	controller *controllers.UserPokemonController,
	supabaseClient *supabase.Client,
) {
	r.With(
		middlewares.SupabaseAuthMiddleware(supabaseClient),
		middlewares.UserRoleMiddleware("client"),
	).Get("/pokedex", controller.GetUserPokedex)

	r.With(
		middlewares.SupabaseAuthMiddleware(supabaseClient),
		middlewares.UserRoleMiddleware("client"),
	).Get("/pokedex/{pokemonId}", controller.GetUserPokemonDetails)
}
