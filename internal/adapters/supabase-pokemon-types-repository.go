package adapters

import (
	"github.com/nedpals/supabase-go"

	"github.com/afmireski/garchop-api/internal/entities"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type SupabasePokemonTypesRepository struct  {
	client *supabase.Client
}

func NewSupabasePokemonTypesRepository(client *supabase.Client) *SupabasePokemonTypesRepository {
	return &SupabasePokemonTypesRepository{
		client: client,
	}
}

func (r *SupabasePokemonTypesRepository) Create(input myTypes.CreatePokemonTypeInput) (*entities.PokemonType, error) {
	var supabaseData myTypes.AnyMap

	err := r.client.DB.From("types").Insert(input).Execute(&supabaseData); if err != nil {
		return nil, err
	}

	return entities.NewPokemonType(
		supabaseData["id"].(string),
		supabaseData["reference_id"].(uint),
		supabaseData["name"].(string),
	), nil
}

func (r *SupabasePokemonTypesRepository) FindByName(name string) (*entities.PokemonType, error) {
	var supabaseData []map[string]interface{}

	err := r.client.DB.From("types").Select("*").Execute(&supabaseData); if err != nil {
		return nil, err
	}

	if len(supabaseData) == 0 {
		return nil, nil
	}

	return entities.NewPokemonType(
		supabaseData[0]["id"].(string),
		supabaseData[0]["reference_id"].(uint),
		supabaseData[0]["name"].(string),
	), nil
}
