package services

import (
	"regexp"
	"time"

	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/validators"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	myTypes "github.com/afmireski/garchop-api/internal/types"
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

func (s *UsersService) UpdateClient(id string, input ports.UpdateUserInput) *customErrors.InternalError {
	if !validators.IsValidUuid(id) {
		return customErrors.NewInternalError("invalid id", 400, []string{})
	}

	if len(input.Name) > 0 && !validators.IsValidName(input.Name, 3, 200) {
		return customErrors.NewInternalError("invalid name", 400, []string{})
	}

	if len(input.Email) > 0 && !validators.IsValidEmail(input.Email) {
		return customErrors.NewInternalError("invalid email", 400, []string{})
	}

	if len(input.Phone) > 0 && !validators.IsPhoneNumber(input.Phone) {
		return customErrors.NewInternalError("invalid phone", 400, []string{})
	}

	return nil
}

func (s *UsersService) DeleteClient(id string) *customErrors.InternalError {
	if !validators.IsValidUuid(id) {
		return customErrors.NewInternalError("invalid uuid", 400, []string{})
	}

	data := myTypes.AnyMap{
		"deleted_at": time.Now(),
		"updated_at": time.Now(),
	}

	where := myTypes.Where{
		"deleted_at": map[string]string{"is": "null"},
	}

	updatedData, err := s.repository.Update(id, data, where)
	if err != nil || updatedData == nil {
		return customErrors.NewInternalError("a failure occurred when try to delete a user", 500, []string{})
	}

	return nil
}
