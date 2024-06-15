package modules

import (
	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/services"
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/nedpals/supabase-go"
)

type TiersModule struct {
	Repository ports.TiersRepositoryPort
	Service    *services.TiersService
	Controller *controllers.TiersController
}

func NewTiersModule(
	supabaseClient *supabase.Client,
) *TiersModule {
	repository := adapters.NewSupabaseTiersRepository(supabaseClient)

	service := services.NewTiersService(repository)

	controller := controllers.NewTiersController(service)

	return &TiersModule{
		Repository: repository,
		Service:    service,
		Controller: controller,
	}
}