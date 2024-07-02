package modules

import (
	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/services"
	"github.com/nedpals/supabase-go"
)

type UsersStatsModule struct {
	Repository ports.UserStatsRepository
	Service    *services.UsersStatsService
}

func NewUsersStatsModule(
	supabaseClient *supabase.Client,
	tiersService *services.TiersService,

) *UsersStatsModule {
	repository := adapters.NewSupabaseUserStatsRepository(supabaseClient)

	service := services.NewUsersStatsService(repository, tiersService)

	return &UsersStatsModule{
		Repository: repository,
		Service:    service,
	}
}