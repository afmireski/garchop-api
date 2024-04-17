package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/services"
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/afmireski/garchop-api/internal/web/routers"
	"github.com/go-chi/chi/v5"
	"github.com/patrickmn/go-cache"

	supabase "github.com/nedpals/supabase-go"
)

func main() {
	supabaseClient := setupSupabase()
	hashHelper := adapters.NewBcryptHashHelper()
	memCache := cache.New(10*time.Minute, 30*time.Minute)

	usersController := setupUsersModule(supabaseClient, hashHelper)

	authController := setupAuthModule(supabaseClient)

	pokemonController := setupPokemonModule(supabaseClient, memCache)

	tiersController := setupTiersModule(supabaseClient)

	r := chi.NewRouter()
	routers.SetupUsersRouter(r, usersController)
	routers.SetupAuthRouter(r, authController)
	routers.SetupPokemonRouter(r, pokemonController)
	routers.SetupTiersRouter(r, tiersController)

	fmt.Println("API is running...")
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	http.ListenAndServe(port, r)
}

func setupSupabase() *supabase.Client {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	return supabase.CreateClient(supabaseUrl, supabaseKey)
}

func setupUsersModule(supabaseClient *supabase.Client, hashHelper ports.HashHelperPort) *controllers.UsersController {
	usersRepository := adapters.NewSupabaseUsersRepository(supabaseClient)
	usersService := services.NewUsersService(usersRepository, hashHelper)
	usersController := controllers.NewUsersController(usersService)

	return usersController
}

func setupAuthModule(supabaseClient *supabase.Client) *controllers.AuthController {
	authAdapter := adapters.NewSupabaseAuthenticator(supabaseClient)
	authService := services.NewAuthService(authAdapter)
	return controllers.NewAuthController(authService)
}

func setupPokemonModule(supabaseClient *supabase.Client, cache *cache.Cache) *controllers.PokemonController {
	pokemonRepository := adapters.NewSupabasePokemonRepository(supabaseClient)
	typeRepository := adapters.NewSupabasePokemonTypesRepository(supabaseClient)
	pokemonService := services.NewPokemonService(pokemonRepository, typeRepository, cache)
	return controllers.NewPokemonController(pokemonService)
}

func setupTiersModule(supabaseClient *supabase.Client) *controllers.TiersController {
	tiersRepository := adapters.NewSupabaseTiersRepository(supabaseClient)
	tiersService := services.NewTiersService(tiersRepository)
	return controllers.NewTiersController(tiersService)
}
