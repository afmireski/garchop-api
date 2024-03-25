package adapters

import (
	"errors"
	"github.com/afmireski/garchop-api/internal/ports"
	supabase "github.com/nedpals/supabase-go"
)

type SupabaseUsersRepository struct {
	client *supabase.Client
}

func NewSupabaseUsersRepository(client *supabase.Client) *SupabaseUsersRepository {
	return &SupabaseUsersRepository{
		client: client,
	}
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

func (r *SupabaseUsersRepository) FindById(id string) (interface{}, error) {
	return nil, errors.New("not implemented")
}

func (r *SupabaseUsersRepository) Update(id string, input interface{}) (interface{}, error) {
	return nil, errors.New("not implemented")
}

func (r *SupabaseUsersRepository) Delete(id string) error {
	return errors.New("not implemented")
}
