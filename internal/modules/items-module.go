package modules

import (
	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/nedpals/supabase-go"
)

type ItemsModule struct {
	Repository ports.ItemsRepositoryPort
}

func NewItemsModule(supabaseClient *supabase.Client) *ItemsModule {
	repository := adapters.NewSupabaseItemsRepository(supabaseClient)

	return &ItemsModule{
		Repository: repository,
	}
}