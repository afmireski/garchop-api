package services

import (
	"fmt"

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
func (s *AuthService) Login(input myTypes.LoginInput) *customErrors.InternalError {

	if inputErr := validateLoginInput(input); inputErr != nil {
		return inputErr
	}

	response, err := s.authenticator.ValidateCredentials(input.Email, input.Password); if err != nil {
		return customErrors.NewInternalError("invalid credentials", 500, []string{})
	}
	
	fmt.Println(response)

	return nil
}