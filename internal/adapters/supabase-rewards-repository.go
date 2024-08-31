package adapters

import (
	"encoding/json"
	"strings"

	"github.com/afmireski/garchop-api/internal/models"
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

func (r *SupabaseRewardsRepository) Create(input myTypes.CreateRewardInput) (string, error) {
	var supabaseData []myTypes.AnyMap

	err := r.client.DB.From("rewards").Insert(input).Execute(&supabaseData)
	
	if err != nil {
		return "", err
	}

	return "", nil
}

func (r *SupabaseRewardsRepository) FindAll(where myTypes.Where, pagination myTypes.Pagination) ([]models.RewardModel, error) {
	var supabaseData []myTypes.AnyMap

	query := r.client.DB.From("rewards").Select("*", "tiers(*)").LimitWithOffset(pagination.Limit, pagination.Offset).Is("deleted_at", "null")

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

func (r *SupabaseRewardsRepository) FindById(id string, where myTypes.Where) (*models.RewardModel, error) {
	var supabaseData myTypes.AnyMap

	query := r.client.DB.From("rewards").Select("*", "tiers(*)").Is("deleted_at", "null")

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

	return r.serializeToModel(supabaseData)
}

func (r *SupabaseRewardsRepository) Update(id string, data myTypes.AnyMap, where myTypes.Where) (*models.RewardModel, error) {
	var supabaseData []myTypes.AnyMap

	query := r.client.DB.From("rewards").Update(data).Eq("id", id).Is("deleted_at", "null");

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
	} else if len(supabaseData) == 0 {
		return nil, nil
	}

	

	return r.serializeToModel(supabaseData[0])
}