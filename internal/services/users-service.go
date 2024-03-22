package services

import (
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
	if (!validators.IsValidEmail(input.Email)) {
		return customErrors.NewInternalError("invalid email", 400, []string{})
	} else if (!validators.IsValidName(input.Name, 3, 200)) {
		return customErrors.NewInternalError("invalid name", 400, []string{})
	} else if (input.Phone == "") {
		return customErrors.NewInternalError("invalid phone", 400, []string{})
	} else if (validators.IsValidateAge(input.BirthDate, 18)) {
		customErrors.NewInternalError("invalid birth date", 400, []string{})
	} 

	return nil
}

func (s *UsersService) NewUser(input ports.CreateUserInput) *customErrors.InternalError {

	if inputErr := validateNewUserInput(input); inputErr != nil {
		return inputErr
	}

	_, err := s.repository.Create(input);

	if err != nil {
		return customErrors.NewInternalError("a failure occurred when try to create a new user", 500, []string{})
	}

	return nil;
}