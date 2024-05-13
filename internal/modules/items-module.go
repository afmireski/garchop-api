package modules

import (
	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/services"
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/nedpals/supabase-go"
)

type ItemsModule struct {
	Repository ports.ItemsRepositoryPort
	Service    *services.ItemsService
	Controller *controllers.ItemController
}

func NewItemsModule(supabaseClient *supabase.Client) *ItemsModule {

	repository := adapters.NewSupabaseItemsRepository(supabaseClient)
	cartsRepository := adapters.NewSupabaseCartsRepository(supabaseClient)

	service := services.NewItemsService(repository, cartsRepository)

	controller := controllers.NewItemController(service)

	return &ItemsModule{
		Repository: repository,
		Service:    service,
		Controller: controller,
	}
}
