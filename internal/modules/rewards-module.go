package modules

import (
	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/services"
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/nedpals/supabase-go"
)

type RewardsModule struct {
	Repository ports.RewardsRepositoryPort
	Service *services.RewardsService
	Controller *controllers.RewardsController
}

func NewRewardsModule(
	supabaseClient *supabase.Client,
	userRewardsRepository ports.UserRewardsRepositoryPort,
	userPokemonsRepository ports.UserPokemonRepositoryPort,
) *RewardsModule {
	repository := adapters.NewSupabaseRewardsRepository(supabaseClient)
	service := services.NewRewardsService(repository, userRewardsRepository, userPokemonsRepository)
	controller := controllers.NewRewardsController(service)

	return &RewardsModule{
		Repository: repository,
		Service: service,
		Controller: controller,
	}
}