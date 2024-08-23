package services

import (
	"github.com/afmireski/garchop-api/internal/entities"
	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/validators"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type UserPokemonsService struct {
	userPokemonRepository ports.UserPokemonRepositoryPort
}

func NewUserPokemonsService(userPokemonRepository ports.UserPokemonRepositoryPort) *UserPokemonsService {
	return &UserPokemonsService{
		userPokemonRepository: userPokemonRepository,
	}
}

func (s* UserPokemonsService) validateGetUserPokedexInput(input myTypes.GetPokedexInput)  *customErrors.InternalError {
	if (!validators.IsValidUuid(input.UserId)) {
		return customErrors.NewInternalError("invalid user_id", 400, []string{"the user_id must be a valid uuid"})
	}

	return nil
}

func (s* UserPokemonsService) GetUserPokedex(input myTypes.GetPokedexInput) ([]entities.UserPokemon, *customErrors.InternalError) {
	if inputErr := s.validateGetUserPokedexInput(input); inputErr != nil {
		return nil, inputErr
	}

	response, err := s.userPokemonRepository.FindAll(myTypes.Where{
		"user_id": map[string]string{"eq": input.UserId},
	}); if err != nil {
		return nil, customErrors.NewInternalError("failed on get the user pokemons", 500, []string{err.Error()})
	}

	return entities.BuilManyUserPokemonFromModel(response), nil
}