package adapters

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

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

func (r *SupabaseUsersRepository) serializeToModel(supabaseData myTypes.AnyMap) (*models.UserModel, error) {
	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		return nil, err
	}

	var modelData models.UserModel
	err = json.Unmarshal(jsonData, &modelData)
	if err != nil {
		return nil, err
	}

	return &modelData, nil
}

func (r *SupabaseUsersRepository) serializeManyToModel(supabaseData []myTypes.AnyMap) ([]models.UserModel, error) {
	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		return nil, err
	}

	var modelsData []models.UserModel
	err = json.Unmarshal(jsonData, &modelsData)
	if err != nil {
		return nil, err
	}

	return modelsData, nil
}

type CreateInput struct {
	Name      string                   `json:"name"`
	Email     string                   `json:"email"`
	Phone     string                   `json:"phone"`
	Password  string                   `json:"password"`
	BirthDate *time.Time               `json:"birth_date"`
	Role      models.UserModelRoleEnum `json:"role"`
}

func (r *SupabaseUsersRepository) Create(input ports.CreateUserInput) (string, error) {
	var supabaseData []map[string]string

	data := CreateInput{
		Name:      input.Name,
		Email:     input.Email,
		Phone:     input.Phone,
		Password:  input.Password,
		BirthDate: input.BirthDate,
		Role:      input.Role,
	}

	err := r.client.DB.From("users").Insert(data).Execute(&supabaseData)
	if err != nil {
		return "", err
	}

	// SignUp the user into supabase auth table
	_, signUpErr := r.client.Auth.SignUp(context.Background(), supabase.UserCredentials{
		Email:    input.Email,
		Password: input.PlainPassword,
	})
	if signUpErr != nil {
		return "", err
	}

	return supabaseData[0]["id"], nil
}

func (r *SupabaseUsersRepository) FindById(id string, where myTypes.Where) (*models.UserModel, error) {
	var supabaseData map[string]interface{}

	query := r.client.DB.From("users").Select("*").Single().Eq("id", id)

	if len(where) > 0 {
		for column, filter := range where {
			for operator, criteria := range filter {
				query = query.Filter(column, operator, criteria)
			}
		}
	}

	err := query.Execute(&supabaseData)
	if err != nil {
		if strings.Contains(err.Error(), "PGRST116") { // resource not found
			return nil, nil
		}

		return nil, err
	}

	return r.serializeToModel(supabaseData)
}

func (r *SupabaseUsersRepository) FindAll(where myTypes.Where) ([]models.UserModel, error) {
	var supabaseData []myTypes.AnyMap

	query := r.client.DB.From("users").Select("*").Is("deleted_at", "null")

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

	if len(supabaseData) == 0 {
		return nil, nil
	}

	return r.serializeManyToModel(supabaseData)
}

func (r *SupabaseUsersRepository) Update(id string, input myTypes.AnyMap, where myTypes.Where) (*models.UserModel, error) {
	var supabaseData []myTypes.AnyMap
	query := r.client.DB.From("users").Update(input).Eq("id", id)
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

	if len(supabaseData) == 0 {
		return nil, nil
	}

	result, err := r.serializeManyToModel(supabaseData)
	if err != nil {
		return nil, err
	}

	return &result[0], nil
}

func (r *SupabaseUsersRepository) Delete(id string) error {
	return errors.New("not implemented")
}
