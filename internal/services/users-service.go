package services

import (
	"regexp"
	"time"

	"github.com/afmireski/garchop-api/internal/entities"
	"github.com/afmireski/garchop-api/internal/models"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/validators"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type UsersService struct {
	repository          ports.UserRepositoryPort
	userStatsRepository ports.UserStatsRepository
	hashHelper          ports.HashHelperPort
}

func NewUsersService(repository ports.UserRepositoryPort, userStatsRepository ports.UserStatsRepository, hashHelper ports.HashHelperPort) *UsersService {
	return &UsersService{
		repository,
		userStatsRepository,
		hashHelper,
	}
}

func validateNewUserInput(input myTypes.NewUserInput) *customErrors.InternalError {
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
	if !validators.IsValidPassword(input.Password) {
		return customErrors.NewInternalError("invalid password", 400, []string{})
	}
	if input.Password != input.ConfirmPassword {
		return customErrors.NewInternalError("passwords do not match", 400, []string{})
	}

	return nil
}

func mapToUpdateClientOutput(input *models.UserModel) *myTypes.UpdateClientOutput {
	if input == nil {
		return nil
	}

	res := myTypes.UpdateClientOutput{
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}

	return &res
}

func (s *UsersService) NewClient(input myTypes.NewUserInput) *customErrors.InternalError {
	if inputErr := validateNewUserInput(input); inputErr != nil {
		return inputErr
	}

	// Remove special characters except '+' from phone
	re := regexp.MustCompile(`[^+\d]`)
	input.Phone = re.ReplaceAllString(input.Phone, "")

	hash, _ := s.hashHelper.GenerateHash(input.Password, 10)

	data := ports.CreateUserInput{
		Name:          input.Name,
		Email:         input.Email,
		Phone:         input.Phone,
		Password:      hash,
		PlainPassword: input.Password,
		BirthDate:     input.BirthDate,
		Role:          models.Client,
	}

	clientId, err := s.repository.Create(data)

	if err != nil {
		return customErrors.NewInternalError("a failure occurred when try to create a new client", 500, []string{err.Error()})
	}

	_, statsErr := s.userStatsRepository.Create(myTypes.CreateUserStatsInput{
		UserId: clientId,
		TierId: 1,
	})
	if statsErr != nil {
		return customErrors.NewInternalError("a failure occurred when try to create client status", 500, []string{statsErr.Error()})
	}

	return nil
}

func (s *UsersService) NewAdmin(input myTypes.NewUserInput) *customErrors.InternalError {
	if inputErr := validateNewUserInput(input); inputErr != nil {
		return inputErr
	}

	// Remove special characters except '+' from phone
	re := regexp.MustCompile(`[^+\d]`)
	input.Phone = re.ReplaceAllString(input.Phone, "")

	hash, _ := s.hashHelper.GenerateHash(input.Password, 10)

	data := ports.CreateUserInput{
		Name:          input.Name,
		Email:         input.Email,
		Phone:         input.Phone,
		Password:      hash,
		PlainPassword: input.Password,
		BirthDate:     input.BirthDate,
		Role:          models.Admin,
	}

	_, err := s.repository.Create(data)

	if err != nil {
		return customErrors.NewInternalError("a failure occurred when try to create a new administrator", 500, []string{err.Error()})
	}

	return nil
}

func (s *UsersService) UpdateClient(id string, input myTypes.UpdateUserInput) (*myTypes.UpdateClientOutput, *customErrors.InternalError) {
	data := myTypes.AnyMap{}

	if !validators.IsValidUuid(id) {
		return nil, customErrors.NewInternalError("invalid id", 400, []string{})
	}

	if len(input.Name) > 0 {
		if !validators.IsValidName(input.Name, 3, 200) {
			return nil, customErrors.NewInternalError("invalid name", 400, []string{})
		}
		data["name"] = input.Name
	}

	if len(input.Email) > 0 {
		if !validators.IsValidEmail(input.Email) {
			return nil, customErrors.NewInternalError("invalid email", 400, []string{})
		}
		data["email"] = input.Email
	}

	if len(input.Phone) > 0 {
		if !validators.IsPhoneNumber(input.Phone) {
			return nil, customErrors.NewInternalError("invalid phone", 400, []string{})
		}
		data["phone"] = input.Phone
	}

	where := myTypes.Where{
		"deleted_at": map[string]string{"is": "null"},
	}
	updatedUser, err := s.repository.Update(id, data, where)

	if err != nil {
		return nil, customErrors.NewInternalError("a failure occurred when trying to update a user", 500, []string{})
	} else if updatedUser == nil {
		return nil, customErrors.NewInternalError("no user found to update", 404, []string{})
	}

	return mapToUpdateClientOutput(updatedUser), nil
}

func (s *UsersService) GetUserById(id string) (*entities.User, *customErrors.InternalError) {
	if !validators.IsValidUuid(id) {
		return nil, customErrors.NewInternalError("invalid id", 400, []string{})
	}

	where := myTypes.Where{
		"deleted_at": map[string]string{"is": "null"},
	}

	response, err := s.repository.FindById(id, where)

	if err != nil {
		return nil, customErrors.NewInternalError("a failure occurred when try to retrieve a new user", 500, []string{err.Error()})
	} else if response == nil {
		return nil, customErrors.NewInternalError("user not found", 404, []string{})
	}

	return entities.NewUser(response.Id, response.Name, response.Email, response.Phone, response.BirthDate, string(response.Role)), nil
}

func (s *UsersService) GetUsers(where myTypes.Where) ([]entities.User, *customErrors.InternalError) {
	repositoryData, err := s.repository.FindAll(where)

	if err != nil {
		return nil, customErrors.NewInternalError("a failure occurred when trying to retrieve users", 500, []string{})
	}

	return entities.BuildManyUserFromModel(repositoryData), nil
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
