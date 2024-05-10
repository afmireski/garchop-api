package adapters

import (
	"encoding/json"

	"github.com/afmireski/garchop-api/internal/models"
	supabase "github.com/nedpals/supabase-go"

	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type SupabaseStocksRepository struct {
	client *supabase.Client
}

func NewSupabaseStocksRepository(client *supabase.Client) *SupabaseStocksRepository {
	return &SupabaseStocksRepository{
		client: client,
	}
}

func (r *SupabaseStocksRepository) serializeToModel(supabaseData myTypes.AnyMap) (*models.StockModel, error) {
	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		return nil, err
	}

	var modelData models.StockModel
	err = json.Unmarshal(jsonData, &modelData)
	if err != nil {
		return nil, err
	}

	return &modelData, nil
}

func (r *SupabaseStocksRepository) Update(id string, input myTypes.AnyMap, where myTypes.Where) (*models.StockModel, error) {
	var supabaseData []myTypes.AnyMap

	query := r.client.DB.From("stocks").Update(input).Eq("pokemon_id", id);

	if len(where) > 0 {
		for column, filter := range where {
			for operator, criteria := range filter {
				query = query.Filter(column, operator, criteria)
			}
		}
	}

	err := query.Execute(&supabaseData); if err != nil {
		return nil, err
	}

	return r.serializeToModel(supabaseData[0])
}


