package modules

import (
	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/services"
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/nedpals/supabase-go"
)

type UsersModule struct {
	Repository ports.UserRepositoryPort
	Service    *services.UsersService
	Controller *controllers.UsersController
}

func NewUsersModule(
	supabaseClient *supabase.Client,
	userStatsRepository ports.UserStatsRepository,
	hashHelper ports.HashHelperPort,
	authService *services.AuthService,
) *UsersModule {
	repository := adapters.NewSupabaseUsersRepository(supabaseClient)
	service := services.NewUsersService(repository, userStatsRepository, hashHelper, authService)
	controller := controllers.NewUsersController(service)

	return &UsersModule{
		Repository: repository,
		Service:    service,
		Controller: controller,
	}
}
