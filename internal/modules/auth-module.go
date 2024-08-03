package modules

import (
	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/services"
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/nedpals/supabase-go"
)

type AuthModule struct {
	Repository ports.AuthenticatorPort
	Service    *services.AuthService
	Controller *controllers.AuthController
}

func NewAuthModule(supabaseClient *supabase.Client) *AuthModule {

	repository := adapters.NewSupabaseAuthenticator(supabaseClient)
	authService := services.NewAuthService(repository)
	controller := controllers.NewAuthController(authService)

	return &AuthModule{
		Repository: repository,
		Service:    authService,
		Controller: controller,
	}
}
