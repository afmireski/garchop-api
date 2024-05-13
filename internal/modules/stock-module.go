package modules

import (
	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/nedpals/supabase-go"
)

type StockModule struct {
	Repository ports.StockRepositoryPort
}

func NewStockModule(supabaseClient *supabase.Client) *StockModule {
	return &StockModule{
		Repository: adapters.NewSupabaseStocksRepository(supabaseClient),
	}
}