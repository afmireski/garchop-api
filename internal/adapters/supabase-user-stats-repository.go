package adapters

import (
	"encoding/json"
	"strings"

	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
	"github.com/nedpals/supabase-go"
)

type SupabaseUserStatsRepository struct {
	client *supabase.Client
}

func NewSupabaseUserStatsRepository(client *supabase.Client) *SupabaseUserStatsRepository {
	return &SupabaseUserStatsRepository{
		client: client,
	}
}

func (s *SupabaseUserStatsRepository) serializeSupabaseDataToModel(supabaseData myTypes.AnyMap) (*models.UserStatsModel, error) {
	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		return nil, err
	}

	var modelData models.UserStatsModel
	err = json.Unmarshal(jsonData, &modelData)
	if err != nil {
		return nil, err
	}

	return &modelData, nil
}

func (s *SupabaseUserStatsRepository) Create(input myTypes.CreateUserStatsInput) (string, error) {
	var supabaseData []myTypes.AnyMap

	err := s.client.DB.From("user_stats").Insert(input).Execute(&supabaseData); if err != nil {
		return "", err
	}

	return supabaseData[0]["user_id"].(string), nil
}

func (s *SupabaseUserStatsRepository) FindById(id string, where myTypes.Where) (*models.UserStatsModel, error) {
	var supabaseData myTypes.AnyMap

	query := s.client.DB.From("user_stats").Select("*", "tiers(*)").Single().Eq("user_id", id)
	
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

	return s.serializeSupabaseDataToModel(supabaseData)
}

func (s *SupabaseUserStatsRepository) Update(id string, input myTypes.AnyMap, where myTypes.Where) (*models.UserStatsModel, error) {
	panic("implement me")
}

func (s *SupabaseUserStatsRepository) Delete(id string, where myTypes.Where) error {
	panic("implement me")
}
