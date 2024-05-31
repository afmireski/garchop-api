package services

import (
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/validators"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type AuthService struct {
	authenticator ports.AuthenticatorPort
}

func NewAuthService(authenticator ports.AuthenticatorPort) *AuthService {
	return &AuthService{
		authenticator,
	}
}

func validateLoginInput(input myTypes.LoginInput) *customErrors.InternalError {
	if !validators.IsValidEmail(input.Email) {
		return customErrors.NewInternalError("invalid email", 400, []string{})
	}
	if !validators.IsValidPassword(input.Password) {
		return customErrors.NewInternalError("invalid password", 400, []string{})
	}
	return nil
}
func (s *AuthService) Login(input myTypes.LoginInput) (*myTypes.LoginOutput, *customErrors.InternalError) {

	if inputErr := validateLoginInput(input); inputErr != nil {
		return nil, inputErr
	}

	response, err := s.authenticator.ValidateCredentials(input.Email, input.Password); if err != nil {
		return nil, customErrors.NewInternalError("invalid credentials", 500, []string{err.Error()})
	}

	return response, nil
}


func (s *AuthService) Logout(token string) *customErrors.InternalError {
	err := s.authenticator.RevogueCredentials(token); if err != nil {
		return customErrors.NewInternalError("failed on sign out", 500, []string{err.Error()})
	}

	return nil;
}