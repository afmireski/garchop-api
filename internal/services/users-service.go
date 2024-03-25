package services

import (
	"regexp"

	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/validators"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
)

type UsersService struct {
	repository ports.UserRepositoryPort
}

func NewUsersService(repository ports.UserRepositoryPort) *UsersService {
	return &UsersService{
		repository,
	}
}

func validateNewUserInput(input ports.CreateUserInput) *customErrors.InternalError {
	if !validators.IsValidEmail(input.Email) {
		return customErrors.NewInternalError("invalid email", 400, []string{})
	} 
	if !validators.IsValidName(input.Name, 3, 200) {
		return customErrors.NewInternalError("invalid name", 400, []string{})
	}  
	if !validators.IsPhoneNumber(input.Phone) {
		return customErrors.NewInternalError("invalid phone", 400, []string{})
	} 
	if !validators.IsValidAge(input.BirthDate, 18) {
		return customErrors.NewInternalError("invalid birth date", 400, []string{})
	}

	return nil
}

func (s *UsersService) NewUser(input ports.CreateUserInput) *customErrors.InternalError {
	if inputErr := validateNewUserInput(input); inputErr != nil {
		return inputErr
	}

	// Remove special characters except '+' from phone
	re := regexp.MustCompile(`[^+\d]`)
	input.Phone = re.ReplaceAllString(input.Phone, "")

	_, err := s.repository.Create(input)

	if err != nil {
		return customErrors.NewInternalError("a failure occurred when try to create a new user", 500, []string{err.Error()})
	}

	return nil
}
