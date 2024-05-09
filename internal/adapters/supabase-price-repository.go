package adapters

import (
	"encoding/json"
	"strings"

	"github.com/afmireski/garchop-api/internal/models"
	supabase "github.com/nedpals/supabase-go"

	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type SupabasePriceRepository struct {
	client *supabase.Client
}

func NewSupabasePriceRepository(client *supabase.Client) *SupabasePriceRepository {
	return &SupabasePriceRepository{
		client: client,
	}
}

func (r *SupabasePriceRepository) serializeToModel(supabaseData myTypes.AnyMap) (*models.PriceModel, error) {
	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		return nil, err
	}

	var modelData models.PriceModel
	err = json.Unmarshal(jsonData, &modelData)
	if err != nil {
		return nil, err
	}

	return &modelData, nil
}

func (r *SupabasePriceRepository) FindCurrentPrice(pokemonId string) (*models.PriceModel, error) {
	var supabaseData myTypes.AnyMap

	err := r.client.DB.From("prices").Select("*", "pokemons (*, stocks (*))").OrderBy("created_at", "desc").Single().Eq("pokemon_id", pokemonId).Is("deleted_at", "null").Execute(&supabaseData); if err != nil {
		if strings.Contains(err.Error(), "PGRST116") { // resource not found
			return nil, nil
		}

		return nil, err
	}

	return r.serializeToModel(supabaseData)
}
