package ports

import myTypes "github.com/afmireski/garchop-api/internal/types"


type AuthenticatorPort interface {
	ValidateCredentials(email string, password string) (*myTypes.LoginOutput, error)
}