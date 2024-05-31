package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/modules"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/services"
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/afmireski/garchop-api/internal/web/routers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	stockModules := modules.NewStockModule(supabaseClient)

	pricesModules := modules.NewPricesModule(supabaseClient)

	itemsModule := modules.NewItemsModule(supabaseClient)

	cartsModule := modules.NewCartsModule(supabaseClient, itemsModule.Repository, pricesModules.Repository, stockModules.Repository)

	purchasesModule := modules.NewPurchasesModule(supabaseClient, cartsModule.Repository, itemsModule.Repository)

	r := chi.NewRouter()
	enableCors(r)
	r.Use(middleware.AllowContentType("application/json"))
	routers.SetupUsersRouter(r, usersController, supabaseClient)
	routers.SetupAuthRouter(r, authController, supabaseClient)
	routers.SetupPokemonRouter(r, pokemonController, supabaseClient)
	routers.SetupTiersRouter(r, tiersController, supabaseClient)
	routers.SetupCartsRouter(r, cartsModule.Controller, supabaseClient)
	routers.SetupItemsRouter(r, itemsModule.Controller, supabaseClient)
	routers.SetupPurchasesRouter(r, purchasesModule.Controller, supabaseClient)

	fmt.Println("API is running...")
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	http.ListenAndServe(port, r)
}

func enableCors(r *chi.Mux) {
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
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

