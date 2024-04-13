package adapters

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/afmireski/garchop-api/internal/models"

	myTypes "github.com/afmireski/garchop-api/internal/types"

	supabase "github.com/nedpals/supabase-go"
)

type SupabaseTiersRepository struct {
	client *supabase.Client
}

func NewSupabaseTiersRepository(client *supabase.Client) *SupabaseTiersRepository {
	return &SupabaseTiersRepository{
		client: client,
	}
}

func serializeSupabaseDataToModel(supabaseData myTypes.AnyMap) (*models.TierModel, error) {
	jsonData, err := json.Marshal(supabaseData); if err != nil {
		return nil, err
	}

	var modelData models.TierModel
	err = json.Unmarshal(jsonData, &modelData); if err != nil {
		return nil, err
	}

	return &modelData, nil
	
}

func serializeManySupabaseDataToModel(supabaseData []myTypes.AnyMap) ([]models.TierModel, error) {
	jsonData, err := json.Marshal(supabaseData); if err != nil {
		return nil, err
	}

	var modelData []models.TierModel
	err = json.Unmarshal(jsonData, &modelData); if err != nil {
		return nil, err
	}

	return modelData, nil
}

func (r *SupabaseTiersRepository) FindAll(where myTypes.Where) ([]models.TierModel, error) {
	var supabaseData []myTypes.AnyMap

	query := r.client.DB.From("tiers").Select("*").Is("deleted_at", "null")

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

	return serializeManySupabaseDataToModel(supabaseData)
}

func (r *SupabaseTiersRepository) FindById(id int, where myTypes.Where) (*models.TierModel, error) {
	var supabaseData myTypes.AnyMap

	query := r.client.DB.From("tiers").Select("*").Single().Eq("id", strconv.Itoa(id));

	if len(where) > 0 {
		for column, filter := range where {
			for operator, criteria := range filter {
				query = query.Filter(column, operator, criteria)
			}
		}
	}
	err := query.Execute(&supabaseData); if err != nil {
		if strings.Contains(err.Error(), "PGRST116") { // resource not found
			return nil, nil
		}
		
		return nil, err
	}

	return serializeSupabaseDataToModel(supabaseData)
}
