package adapters

import (
	"context"

	supabase "github.com/nedpals/supabase-go"

	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type SupabaseAuthenticator struct {
	client *supabase.Client
}

func NewSupabaseAuthenticator(client *supabase.Client) *SupabaseAuthenticator {
	return &SupabaseAuthenticator{
		client: client,
	}
}


func (a *SupabaseAuthenticator) ValidateCredentials(email string, password string) (*myTypes.LoginOutput, error) {
	ctx := context.Background()
	// Valdiate credentials and create a session
	authData, err := a.client.Auth.SignIn(ctx, supabase.UserCredentials{
		Email:    email,
		Password: password,
	})

	return &myTypes.LoginOutput{
		AccessToken: authData.AccessToken,
		RefreshToken: authData.RefreshToken,
	}, err
}