package adapters

import (
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

func (s *SupabaseUserStatsRepository) Create(input myTypes.CreateUserStatsInput) (string, error) {
	var supabaseData []myTypes.AnyMap

	err := s.client.DB.From("user_stats").Insert(input).Execute(&supabaseData); if err != nil {
		return "", err
	}

	return supabaseData[0]["user_id"].(string), nil
}

func (s *SupabaseUserStatsRepository) FindById(id string, where myTypes.Where) (*models.UserStatsModel, error) {
	panic("implement me")
}

func (s *SupabaseUserStatsRepository) Update(id string, input myTypes.AnyMap, where myTypes.Where) (*models.UserStatsModel, error) {
	panic("implement me")
}

func (s *SupabaseUserStatsRepository) Delete(id string, where myTypes.Where) error {
	panic("implement me")
}
