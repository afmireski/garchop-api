package modules

import (
	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/services"
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/nedpals/supabase-go"
)

type PokedexModule struct {
	Repository ports.UserPokemonRepositoryPort
	Service    *services.UserPokemonsService
	Controller *controllers.UserPokemonController	
}

func NewUserPokemonModule(
	supabaseClient *supabase.Client,
) *PokedexModule {
	repository := adapters.NewSupabaseUserPokemonRepository(supabaseClient)
	service := services.NewUserPokemonsService(repository)
	controller := controllers.NewUserPokemonController(service)

	return &PokedexModule{
		Repository: repository,
		Service:    service,
		Controller: controller,
	}
}