package adapters

import (
	"encoding/json"
	"errors"

	"github.com/afmireski/garchop-api/internal/models"
	"github.com/afmireski/garchop-api/internal/ports"
	supabase "github.com/nedpals/supabase-go"

	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type SupabaseUsersRepository struct {
	client *supabase.Client
}

func NewSupabaseUsersRepository(client *supabase.Client) *SupabaseUsersRepository {
	return &SupabaseUsersRepository{
		client: client,
	}
}

func serializeMany(data []map[string]string) ([]models.UserModel, error) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}


	var result []models.UserModel
	json.Unmarshal(jsonData, &result)

	return result, nil
}

func (r *SupabaseUsersRepository) Create(input ports.CreateUserInput) (string, error) {
	var supabaseData []map[string]string

	err := r.client.DB.From("users").Insert(input).Execute(&supabaseData)

	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return supabaseData[0]["id"], nil
}

func (r *SupabaseUsersRepository) FindById(id string) (myTypes.Any, error) {
	return nil, errors.New("not implemented")
}

func (r *SupabaseUsersRepository) Update(id string, input myTypes.AnyMap, where myTypes.Where) (myTypes.Any, error) {
	var supabaseData []map[string]string
	query := r.client.DB.From("users").Update(input).Eq("id", id)
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

	if (len(supabaseData) == 0) {
		return nil, nil
	}

	result, err := serializeMany(supabaseData); if err != nil {
		return nil, err
	}	

	return result[0], nil
}

func (r *SupabaseUsersRepository) Delete(id string) error {
	return errors.New("not implemented")
}
