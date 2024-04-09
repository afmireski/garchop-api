package adapters

import (
	"encoding/json"
	"strings"

	supabase "github.com/nedpals/supabase-go"

	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type SupabasePokemonRepository struct {
	client *supabase.Client
}

func NewSupabasePokemonRepository(client *supabase.Client) *SupabasePokemonRepository {
	return &SupabasePokemonRepository{
		client: client,
	}
}

func serializeToModel(supabaseData myTypes.AnyMap) (*models.PokemonModel, error) {
	jsonData, err := json.Marshal(supabaseData); if err != nil {
		return nil, err
	}

	var modelData models.PokemonModel
	err = json.Unmarshal(jsonData, &modelData); if err != nil {
		return nil, err
	}
	
	return &modelData, nil
}

func (r *SupabasePokemonRepository) Create(input myTypes.CreatePokemonInput) (string, error) {
	var supabaseData []myTypes.AnyMap

	err := r.client.DB.From("pokemons").Insert(input).Execute(&supabaseData); if err != nil {
		return "", err
	}

	return supabaseData[0]["id"].(string), nil
}

type createPriceInput struct {
	PokemonId string `json:"pokemon_id"`
	Value int `json:"value"`
}
type createStockInput struct {
	PokemonId string `json:"pokemon_id"`
	Quantity int `json:"quantity"`
}
type createPokemonTypeInput struct {
	PokemonId string `json:"pokemon_id"`
	TypeId string `json:"type_id"`
}
func (r *SupabasePokemonRepository) Registry(input myTypes.RegistryPokemonInput) (string, error) {
	pokemonId, err := r.Create(input.CreatePokemonInput); if err != nil {
		return "", err
	}

	var supabaseData []myTypes.AnyMap
	
	priceInput := createPriceInput{
		PokemonId: pokemonId,
		Value: input.Price,
	}
	err = r.client.DB.From("prices").Insert(priceInput).Execute(&supabaseData); if err != nil {
		return "", err
	}

	stockInput := createStockInput{
		PokemonId: pokemonId,
		Quantity: input.InitialStock,
	}
	err = r.client.DB.From("stocks").Insert(stockInput).Execute(&supabaseData); if err != nil {
		return "", err
	}

	for _, typeId := range input.Types {
		err = r.client.DB.From("pokemon_types").Insert(createPokemonTypeInput{
			PokemonId: pokemonId,
			TypeId: typeId,
		}).Execute(&supabaseData); if err != nil {
			return "", err
		}; if err != nil {
			return "", err
		}
	}

	return pokemonId, nil
}

func (r *SupabasePokemonRepository) FindById(id string) (*models.PokemonModel, error) {
	var supabaseData myTypes.AnyMap

	err := r.client.DB.From("pokemons").Select("*").Single().Eq("id", id).Execute(&supabaseData); if err != nil {
		if strings.Contains(err.Error(), "PGRST116") { // resource not found
			return nil, nil
		}

		return nil, err
	}

	return serializeToModel(supabaseData)
}

func (r *SupabasePokemonRepository) FindAll(where myTypes.Where) ([]myTypes.Any, error) {
	panic("method not implemented")
}

func (r *SupabasePokemonRepository) Update(id string, input myTypes.AnyMap, where myTypes.Where) (myTypes.Any, error) {
	panic("method not implemented")
}




