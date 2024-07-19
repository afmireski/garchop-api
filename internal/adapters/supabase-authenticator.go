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
	// Validiate credentials and create a session
	authData, err := a.client.Auth.SignIn(ctx, supabase.UserCredentials{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return &myTypes.LoginOutput{
		AccessToken:  authData.AccessToken,
		RefreshToken: authData.RefreshToken,
	}, err
}

func (a *SupabaseAuthenticator) RevogueCredentials(token string) error {
	ctx := context.Background()

	return a.client.Auth.SignOut(ctx, token);
}
