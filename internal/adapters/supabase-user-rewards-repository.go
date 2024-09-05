package adapters

import (
	"encoding/json"

	"github.com/afmireski/garchop-api/internal/models"
	"github.com/nedpals/supabase-go"

	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type SupabaseUserRewardsRepository struct {
	client *supabase.Client
}

func NewSupabaseUserRewardsRepository(client *supabase.Client) *SupabaseUserRewardsRepository {
	return &SupabaseUserRewardsRepository{
		client: client,
	}
}

func (s *SupabaseUserRewardsRepository) serializeSupabaseDataToModel(supabaseData myTypes.AnyMap) (*models.UserRewardModel, error) {
	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		return nil, err
	}

	var modelData models.UserRewardModel
	err = json.Unmarshal(jsonData, &modelData)
	if err != nil {
		return nil, err
	}

	return &modelData, nil
}

func (s *SupabaseUserRewardsRepository) serializeManyToModel(supabaseData []myTypes.AnyMap) ([]models.UserRewardModel, error) {
	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		return nil, err
	}

	var modelData []models.UserRewardModel
	err = json.Unmarshal(jsonData, &modelData)
	if err != nil {
		return nil, err
	}

	return modelData, nil
}

func (s *SupabaseUserRewardsRepository) Create(input myTypes.UserRewardInput) (*models.UserRewardModel, error) {
	var supabaseData []myTypes.AnyMap

	err := s.client.DB.From("users_rewards").Insert(input).Execute(&supabaseData)
	if err != nil {
		return nil, err
	}

	return s.serializeSupabaseDataToModel(supabaseData[0])
}

func (s *SupabaseUserRewardsRepository) FindById(input myTypes.UserRewardInput, where myTypes.Where) (models.UserRewardModel, error) {
	panic("implement me")
}

func (s *SupabaseUserRewardsRepository) FindAll(where myTypes.Where) ([]models.UserRewardModel, error) {
	var supabaseData []myTypes.AnyMap

	query := s.client.DB.From("users_rewards").Select("*").Not().Is("user_id", "null")
	
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

	return s.serializeManyToModel(supabaseData)
}

