package adapters

import (
	"github.com/nedpals/supabase-go"

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

func (r *SupabasePokemonTypesRepository) Create(input myTypes.CreatePokemonTypeInput) (string, error) {
	var supabaseData []map[string]string

	err := r.client.DB.From("types").Insert(input).Execute(&supabaseData); if err != nil {
		return "", err
	}

	return supabaseData[0]["id"], nil
}