package adapters

import (
	"context"
	"encoding/json"
	"errors"
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

func serializeMany(data []map[string]string) ([]models.UserModel, error) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}


	var result []models.UserModel
	json.Unmarshal(jsonData, &result)

	return result, nil
}

type CreateInput struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	BirthDate time.Time `json:"birth_date"`
}
func (r *SupabaseUsersRepository) Create(input ports.CreateUserInput) (string, error) {
	var supabaseData []map[string]string

	data := CreateInput{
		Name: input.Name,
		Email: input.Email,
		Phone: input.Phone,
		Password: input.Password,
		BirthDate: input.BirthDate,
	}	

	err := r.client.DB.From("users").Insert(data).Execute(&supabaseData); if err != nil {
		return "", err
	}

	// SignUp the user into supabase auth table
	_, signUpErr := r.client.Auth.SignUp(context.Background(), supabase.UserCredentials{
		Email: input.Email,
		Password: input.PlainPassword,
	}); if signUpErr != nil {
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
