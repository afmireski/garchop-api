package services

import (
	"fmt"

	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/validators"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
)

type AuthService struct {
	authenticator ports.AuthenticatorPort
}

func NewAuthService(authenticator ports.AuthenticatorPort) *AuthService {
	return &AuthService{
		authenticator,
	}
}

func validateLoginInput(email string, password string) *customErrors.InternalError {
	if !validators.IsValidEmail(email) {
		return customErrors.NewInternalError("invalid email", 400, []string{})
	}
	if !validators.IsValidPassword(password) {
		return customErrors.NewInternalError("invalid password", 400, []string{})
	}
	return nil
}
func (s *AuthService) Login(email string, password string) error {

	if inputErr := validateLoginInput(email, password); inputErr != nil {
		return inputErr
	}

	response, err := s.authenticator.ValidateCredentials(email, password); if err != nil {
		return customErrors.NewInternalError("invalid credentials", 500, []string{})
	}
	
	fmt.Println(response)

	return nil
}