package modules

import (
	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/nedpals/supabase-go"
)

type PricesModule struct {
	Repository ports.PriceRepositoryPort
}

func NewPricesModule(supabaseClient *supabase.Client) *PricesModule {
	repository := adapters.NewSupabasePriceRepository(supabaseClient)

	return &PricesModule{
		Repository: repository,
	}
}