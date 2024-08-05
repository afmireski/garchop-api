package modules

import (
	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/services"
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/nedpals/supabase-go"
)

type PaymentsMethodsModule struct {
	Repository ports.PaymentMethodsRepositoryPort
	Service    *services.PaymentsMethodsService
	Controller *controllers.PaymentsMethodsController
}

func NewPaymentsMethodsModule(supabaseClient *supabase.Client) *PaymentsMethodsModule {
	repository := adapters.NewSupabasePaymentMethodsRepository(supabaseClient)
	service := services.NewPaymentsMethodsService(repository)
	controller := controllers.NewPaymentsMethodsController(service)

	return &PaymentsMethodsModule{
		Repository: repository,
		Service:    service,
		Controller: controller,
	}
}
