package modules

import (
	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/services"
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/nedpals/supabase-go"
)

type CartsModule struct {
	Repository ports.CartsRepositoryPort
	Service    *services.CartsService
	Controller *controllers.CartController
}

func NewCartsModule(
	supabaseClient *supabase.Client,
	itemsRepository ports.ItemsRepositoryPort,
	priceRepository ports.PriceRepositoryPort,
	stockRepository ports.StockRepositoryPort) *CartsModule {
		repository := adapters.NewSupabaseCartsRepository(supabaseClient)

		service := services.NewCartsService(repository, itemsRepository, priceRepository, stockRepository)

		controller := controllers.NewCartController(service)

		return &CartsModule{
			Repository: repository,
			Service:    service,
			Controller: controller,
		}
}
