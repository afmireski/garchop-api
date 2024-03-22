package adapters

import (
	"github.com/afmireski/garchop-api/internal/models"
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
	var result models.UserModel;

	err := r.client.DB.From("users").Insert(input).Execute(&result);

	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return result.Id, nil
}
