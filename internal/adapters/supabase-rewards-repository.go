package adapters

import (
	"encoding/json"

	"github.com/afmireski/garchop-api/internal/models"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/nedpals/supabase-go"

	myTypes "github.com/afmireski/garchop-api/internal/types"
)


type SupabaseRewardsRepository struct {
	client *supabase.Client
}

func NewSupabaseRewardsRepository(client *supabase.Client) *SupabaseRewardsRepository {
	return &SupabaseRewardsRepository{
		client: client,
	}
}

func (r *SupabaseRewardsRepository) serializeToModel(supabaseData myTypes.AnyMap) (*models.RewardModel, error) {
	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		return nil, err
	}

	var modelData models.RewardModel
	err = json.Unmarshal(jsonData, &modelData)
	if err != nil {
		return nil, err
	}

	return &modelData, nil
}

func (r *SupabaseRewardsRepository) serializeManyToModel(supabaseData []myTypes.AnyMap) ([]models.RewardModel, error) {
	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		return nil, err
	}

	var modelData []models.RewardModel
	err = json.Unmarshal(jsonData, &modelData)
	if err != nil {
		return nil, err
	}

	return modelData, nil
}

func (r *SupabaseRewardsRepository) Create(input ports.CreateRewardInput) (string, error) {
	var supabaseData []myTypes.AnyMap

	err := r.client.DB.From("rewards").Insert(input).Execute(&supabaseData)
	
	if err != nil {
		return "", err
	}

	return "", nil
}

func (r *SupabaseRewardsRepository) FindAll(where myTypes.Where) ([]models.RewardModel, error) {
	var supabaseData []myTypes.AnyMap

	query := r.client.DB.From("rewards").Select("*", "tiers(*)").Is("deleted_at", "null")

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

	return r.serializeManyToModel(supabaseData)
}

func (r *SupabaseRewardsRepository) Delete(id string, where myTypes.Where) error {
	panic("implement me")
}