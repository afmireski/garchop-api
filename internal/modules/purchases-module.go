package modules

import (
	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/services"
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/nedpals/supabase-go"
)

type PurchasesModule struct {
	Repository ports.PurchaseRepositoryPort
	Service    *services.PurchasesService
	Controller *controllers.PurchaseController
}

func NewPurchasesModule(
	supabaseClient *supabase.Client,
	cartsRepository ports.CartsRepositoryPort,
	itemsRepository ports.ItemsRepositoryPort,
	userPokemonRepository ports.UserPokemonRepositoryPort,
	) *PurchasesModule {
		repository := adapters.NewSupabasePurchaseRepository(supabaseClient)

		service := services.NewPurchasesService(repository, cartsRepository, itemsRepository, userPokemonRepository)

		controller := controllers.NewPurchaseController(service)

		return &PurchasesModule{
			Repository: repository,
			Service:    service,
			Controller: controller,
		}
}