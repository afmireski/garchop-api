package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/services"
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/afmireski/garchop-api/internal/web/routers"
	"github.com/go-chi/chi/v5"

	supabase "github.com/nedpals/supabase-go"
)

func main() {
	supabaseClient := setupSupabase()

	usersController := setupUsersModule(supabaseClient)

	r := chi.NewRouter()
	routers.SetupUsersRouter(r, usersController)

	fmt.Println("API is running...")
	http.ListenAndServe(":3000", r)
}

func setupSupabase() *supabase.Client {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_API_KEY")
	return supabase.CreateClient(supabaseUrl, supabaseKey)
}

func setupUsersModule(supabaseClient *supabase.Client) *controllers.UsersController {
	usersRepository := adapters.NewSupabaseUsersRepository(supabaseClient)
	usersService := services.NewUsersService(usersRepository)
	usersController := controllers.NewUsersController(usersService)

	return usersController
}
