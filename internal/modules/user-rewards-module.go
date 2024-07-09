package modules

import (
	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/nedpals/supabase-go"
)

type UserRewardsModule struct {
	Repository ports.UserRewardsRepositoryPort
}

func NewUserRewardsModule(supabaseClient *supabase.Client) *UserRewardsModule {
	repository := adapters.NewSupabaseUserRewardsRepository(supabaseClient)
	return &UserRewardsModule{
		Repository: repository,
	}
}