package adapters

import (
	"encoding/json"
	"strings"
	"time"

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
	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		return nil, err
	}

	var modelData models.PokemonModel
	err = json.Unmarshal(jsonData, &modelData)
	if err != nil {
		return nil, err
	}

	return &modelData, nil
}

func serializeManyToModel(supabaseData []myTypes.AnyMap) ([]models.PokemonModel, error) {
	timeLayout := "2006-01-02T15:04:05.999999-07:00"

	for _, d := range supabaseData {
		for key, value := range d {
			if strings.Contains(key, "deleted_at") {
				if strValue, ok := value.(string); ok {
					if len(strValue) == 0 {
						tmp := time.Time{}
						d[key] = tmp.Format(timeLayout)
					}
				}
			}
		}
	}

	jsonData, err := json.Marshal(supabaseData)

	if err != nil {
		return nil, err
	}

	var result []models.PokemonModel
	err = json.Unmarshal(jsonData, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *SupabasePokemonRepository) Create(input myTypes.CreatePokemonInput) (string, error) {
	var supabaseData []myTypes.AnyMap

	err := r.client.DB.From("pokemons").Insert(input).Execute(&supabaseData)
	if err != nil {
		return "", err
	}

	return supabaseData[0]["id"].(string), nil
}

type createPriceInput struct {
	PokemonId string `json:"pokemon_id"`
	Value     int    `json:"value"`
}
type createStockInput struct {
	PokemonId string `json:"pokemon_id"`
	Quantity  int    `json:"quantity"`
}

type updateStockInput struct {
	Quantity  int    `json:"quantity"`	
}

type createPokemonTypeInput struct {
	PokemonId string `json:"pokemon_id"`
	TypeId    string `json:"type_id"`
}

func (r *SupabasePokemonRepository) Registry(input myTypes.RegistryPokemonInput) (string, error) {
	pokemonId, err := r.Create(input.CreatePokemonInput)
	if err != nil {
		return "", err
	}

	var supabaseData []myTypes.AnyMap

	priceInput := createPriceInput{
		PokemonId: pokemonId,
		Value:     input.Price,
	}
	err = r.client.DB.From("prices").Insert(priceInput).Execute(&supabaseData)
	if err != nil {
		return "", err
	}

	stockInput := createStockInput{
		PokemonId: pokemonId,
		Quantity:  input.InitialStock,
	}
	err = r.client.DB.From("stocks").Insert(stockInput).Execute(&supabaseData)
	if err != nil {
		return "", err
	}

	for _, typeId := range input.Types {
		err = r.client.DB.From("pokemon_types").Insert(createPokemonTypeInput{
			PokemonId: pokemonId,
			TypeId:    typeId,
		}).Execute(&supabaseData)
		if err != nil {
			return "", err
		}
		if err != nil {
			return "", err
		}
	}

	return pokemonId, nil
}

func (r *SupabasePokemonRepository) FindById(id string, where myTypes.Where) (*models.PokemonModel, error) {
	var supabaseData myTypes.AnyMap

	query := r.client.DB.From("pokemons").Select("*", "prices (*)", "stocks (*)", "pokemon_types (*, types (*))", "tiers (*)").Single().Eq("id", id)

	if len(where) > 0 {
		for column, filter := range where {
			for operator, criteria := range filter {
				query = query.Filter(column, operator, criteria)
			}
		}
	}

	err := query.Execute(&supabaseData)
	if err != nil {
		if strings.Contains(err.Error(), "PGRST116") { // resource not found
			return nil, nil
		}

		return nil, err
	}

	return serializeToModel(supabaseData)
}

func (r *SupabasePokemonRepository) FindAll(where myTypes.Where) ([]models.PokemonModel, error) {
	var supabaseData []myTypes.AnyMap

	query := r.client.DB.From("pokemons").Select("*", "prices (*)", "stocks!inner(*)", "pokemon_types (*, types (*))", "tiers (*)").Is("deleted_at", "null")

	if len(where) > 0 {
		for column, filter := range where {
			for operator, criteria := range filter {
				query = query.Filter(column, operator, criteria)
			}
		}
	}

	err := query.Execute(&supabaseData)

	if err != nil {
		if strings.Contains(err.Error(), "PGRST116") { // resource not found
			return nil, nil
		}

		return nil, err
	}

	return serializeManyToModel(supabaseData)

}

func (r *SupabasePokemonRepository) Update(id string, input myTypes.AnyMap, where myTypes.Where) (*models.PokemonModel, error) {
	var supabaseData []myTypes.AnyMap

	if _, exists := input["price"]; exists {
		priceInput := createPriceInput{
			PokemonId: id,
			Value:     input["price"].(int),
		}

		delete(input, "price")

		err := r.client.DB.From("prices").Insert(priceInput).Execute(&supabaseData)
		if err != nil {
			return nil, err
		}
	}

	if _, exists := input["stock"]; exists {
		stockInput := updateStockInput{
			Quantity:  input["stock"].(int),
		}

		delete(input, "stock")

		err := r.client.DB.From("stocks").Update(stockInput).Eq("pokemon_id", id).Filter("deleted_at", "is", "null").Execute(&supabaseData)
		if err != nil {
			return nil, err
		}
	}

	query := r.client.DB.From("pokemons").Update(input).Eq("id", id)
	if len(where) > 0 {
		for column, filter := range where {
			for operator, criteria := range filter {
				query = query.Filter(column, operator, criteria)
			}
		}
	}

	err := query.Execute(&supabaseData)
	if err != nil {
		return nil, err
	}

	if len(supabaseData) == 0 {
		return nil, nil
	}

	return serializeToModel(supabaseData[0])
}
