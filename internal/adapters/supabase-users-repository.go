package adapters

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/afmireski/garchop-api/internal/entities"
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

func mapToUserModel(data map[string]interface{}) (*models.UserModel, error) {
	birthDate, _ := time.Parse("2006-01-02", data["birth_date"].(string))

	createdAt, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", data["created_at"].(string))

	updatedAt, _ := time.Parse("2006-01-02T15:04:05.999999999Z07:00", data["updated_at"].(string))

	var deletedAt time.Time
	if deletedAtString, ok := data["deleted_at"].(string); ok {
		deletedAt, _ = time.Parse("2006-01-02T15:04:05.999999999Z07:00", deletedAtString)
	}
	
	var role entities.UserRoleEnum;
	if data["role"] == "client" {
		role = entities.Client
	} else {
		role = entities.Admin
	}

	return models.NewUserModel(
		data["id"].(string),
		data["name"].(string),
		data["email"].(string),
		data["phone"].(string),
		birthDate,
		role,
		createdAt,
		updatedAt,
		deletedAt), nil
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

func (r *SupabaseUsersRepository) FindById(id string) (*models.UserModel, error) {
	var supabaseData map[string]interface{}

	err := r.client.DB.From("users").Select("*").Single().Eq("id", id).Execute(&supabaseData)

	if err != nil {

		if strings.Contains(err.Error(), "PGRST116") { // resource not found
			return nil, nil
		}

		return nil, err
	}

	return mapToUserModel(supabaseData)
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
