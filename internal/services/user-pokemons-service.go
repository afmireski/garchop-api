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
		return nil, customErrors.NewInternalError("failed on get the user pokedex", 500, []string{err.Error()})
	}

	return entities.BuilManyUserPokemonFromModel(response), nil
}

func (s* UserPokemonsService) validateGetUserPokemonInput(input myTypes.GetUserPokemonInput)  *customErrors.InternalError {
	if (!validators.IsValidUuid(input.UserId)) {
		return customErrors.NewInternalError("invalid user_id", 400, []string{"the user_id must be a valid uuid"})
	} else if (!validators.IsValidUuid(input.PokemonId)) {
		return customErrors.NewInternalError("invalid pokemon_id", 400, []string{"the pokemon_id must be a valid uuid"})
	}

	return nil
}	

func (s* UserPokemonsService) GetUserPokemon(input myTypes.GetUserPokemonInput) (entities.UserPokemon, *customErrors.InternalError) {
	if inputErr := s.validateGetUserPokemonInput(input); inputErr != nil {
		return entities.UserPokemon{}, inputErr
	}

	response, err := s.userPokemonRepository.FindById(input.UserId, input.PokemonId, myTypes.Where{}); if err != nil {
		return entities.UserPokemon{}, customErrors.NewInternalError("failed on get the user pokemon", 500, []string{err.Error()})
	}
	
	return entities.BuildUserPokemonFromModel(*response), nil
}
