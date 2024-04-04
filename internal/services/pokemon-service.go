package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/validators"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type PokemonService struct {
	repository ports.PokemonRepositoryPort
}

func NewPokemonService(repository ports.PokemonRepositoryPort) *PokemonService {
	return &PokemonService{
		repository: repository,
	}
}

func validateNewPokemonInput(input myTypes.NewPokemonInput) *customErrors.InternalError {
	if (!validators.IsValidName(input.Name, 1, 100)) {
		return customErrors.NewInternalError("invalid name", 400, []string{"Name cannot be empty and must be between 1 and 100 characters"})
	} else if (!validators.IsValidUuid(input.TierId)) {
		return customErrors.NewInternalError("invalid tier_id", 400, []string{"Tier id must be a valid uuid"})
	} else if (!validators.IsPositiveNumber(input.Price)) {
		return customErrors.NewInternalError("invalid price", 400, []string{"Price must be a positive number"})
	} else if (!validators.IsGreaterThanEqualInt(input.InitialStock, 0)) {
		return customErrors.NewInternalError("invalid initial stock", 400, []string{"Initial stock must be greater than or equal to 0"})
	}

	return nil
}

func obtainPokemonData(pokemonName string) (myTypes.AnyMap, *customErrors.InternalError) {
	url := fmt.Sprintf("%s/pokemon/%s", os.Getenv("POKE_API_URL"), strings.ToLower(pokemonName))
	pokeApiResponse, err := http.Get(url); if err != nil || pokeApiResponse.StatusCode != 200 {
		return nil, customErrors.NewInternalError("a failure occurred when try to obtain pokémon data", 500, []string{err.Error()})
	}

	pokeRawData, err := io.ReadAll(pokeApiResponse.Body)
	pokeApiResponse.Body.Close()

	if err != nil {
		return nil, customErrors.NewInternalError("a failure occurred when try to obtain pokémon data", 500, []string{err.Error()})
	}

	var pokeJson myTypes.AnyMap
	err = json.Unmarshal(pokeRawData, &pokeJson) ; if err != nil {
		return nil, customErrors.NewInternalError("a failure occurred when try to obtain pokémon data", 500, []string{err.Error()})
	}

	return pokeJson, nil
}

func (s* PokemonService) NewPokemon(input myTypes.NewPokemonInput) *customErrors.InternalError {

	if inputErr := validateNewPokemonInput(input); inputErr != nil {
		return inputErr
	}

	pokeData, err := obtainPokemonData(input.Name) ; if err != nil {
		return err
	}

	data := myTypes.CreatePokemonInput{
		ReferenceId: pokeData["id"].(uint),
		TierId:      input.TierId,
		Name: input.Name,
		Weight: pokeData["weight"].(uint),
		Height: pokeData["height"].(uint),
		Experience: pokeData["base_experience"].(uint),
		ImageUrl: (((pokeData["sprites"].(myTypes.AnyMap))["other"].(myTypes.AnyMap))["official-artwork"].(myTypes.AnyMap))["front_default"].(string),
	}

	_, repositoryErr := s.repository.Create(data); if repositoryErr != nil {
		return customErrors.NewInternalError("a failure occurred when try to registry a new pokemon", 500, []string{})
	}

	return nil
}