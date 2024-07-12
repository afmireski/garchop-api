package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/afmireski/garchop-api/internal/adapters"
	"github.com/afmireski/garchop-api/internal/modules"
	"github.com/afmireski/garchop-api/internal/services"
	"github.com/afmireski/garchop-api/internal/web/controllers"
	"github.com/afmireski/garchop-api/internal/web/routers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/patrickmn/go-cache"

	supabase "github.com/nedpals/supabase-go"
)

func main() {
	supabaseClient := setupSupabase()
	hashHelper := adapters.NewBcryptHashHelper()
	memCache := cache.New(10*time.Minute, 30*time.Minute)

	tiersModule := modules.NewTiersModule(supabaseClient)

	userStatsModule := modules.NewUsersStatsModule(supabaseClient, tiersModule.Service)
	
	usersModule := modules.NewUsersModule(supabaseClient, userStatsModule.Repository, hashHelper)

	authController := setupAuthModule(supabaseClient)

	pokemonController := setupPokemonModule(supabaseClient, memCache)


	stockModules := modules.NewStockModule(supabaseClient)

	pricesModules := modules.NewPricesModule(supabaseClient)

	itemsModule := modules.NewItemsModule(supabaseClient)

	cartsModule := modules.NewCartsModule(supabaseClient, itemsModule.Repository, pricesModules.Repository, stockModules.Repository)

	purchasesModule := modules.NewPurchasesModule(supabaseClient, cartsModule.Repository, itemsModule.Repository, userStatsModule.Service)

	r := chi.NewRouter()
	enableCors(r)
	routers.SetupUsersRouter(r, usersModule.Controller)
	routers.SetupAuthRouter(r, authController)
	routers.SetupPokemonRouter(r, pokemonController)
	routers.SetupTiersRouter(r, tiersModule.Controller)
	routers.SetupCartsRouter(r, cartsModule.Controller)
	routers.SetupItemsRouter(r, itemsModule.Controller)
	routers.SetupPurchasesRouter(r, purchasesModule.Controller)

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
