package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/afmireski/garchop-api/internal/ports"
	"github.com/afmireski/garchop-api/internal/validators"
	cache "github.com/patrickmn/go-cache"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type PokemonService struct {
	repository      ports.PokemonRepositoryPort
	typesRepository ports.PokemonTypesRepositoryPort
	cache           *cache.Cache
}

func NewPokemonService(repository ports.PokemonRepositoryPort, typesRepository ports.PokemonTypesRepositoryPort, cache *cache.Cache) *PokemonService {
	return &PokemonService{
		repository:      repository,
		typesRepository: typesRepository,
		cache:           cache,
	}
}

func validateNewPokemonInput(input myTypes.NewPokemonInput) *customErrors.InternalError {
	if !validators.IsValidName(input.Name, 1, 100) {
		return customErrors.NewInternalError("invalid name", 400, []string{"Name cannot be empty and must be between 1 and 100 characters"})
	} else if !validators.IsValidNumericId(input.TierId) {
		return customErrors.NewInternalError("invalid tier_id", 400, []string{"Tier id must be a valid uuid"})
	} else if !validators.IsPositiveNumber(input.Price) {
		return customErrors.NewInternalError("invalid price", 400, []string{"Price must be a positive number"})
	} else if !validators.IsGreaterThanEqualInt(input.InitialStock, 0) {
		return customErrors.NewInternalError("invalid initial stock", 400, []string{"Initial stock must be greater than or equal to 0"})
	}

	return nil
}

func searchPokemonInPokeApi(pokemonName string) (myTypes.AnyMap, *customErrors.InternalError) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	fmt.Println(url)
	pokeApiResponse, err := http.Get(url)
	if err != nil || pokeApiResponse.StatusCode != 200 {
		return nil, customErrors.NewInternalError("a failure occurred when try to obtain pokémon data", 500, []string{err.Error()})
	}

	pokeRawData, err := io.ReadAll(pokeApiResponse.Body)
	pokeApiResponse.Body.Close()

	if err != nil {
		return nil, customErrors.NewInternalError("a failure occurred when try to obtain pokémon data", 500, []string{err.Error()})
	}

	var pokeJson myTypes.AnyMap
	err = json.Unmarshal(pokeRawData, &pokeJson)
	if err != nil {
		return nil, customErrors.NewInternalError("a failure occurred when try to obtain pokémon data", 500, []string{err.Error()})
	}

	return pokeJson, nil
}

func (s *PokemonService) obtainPokemonData(pokemonName string) (myTypes.AnyMap, *customErrors.InternalError) {
	pokeDataKey := fmt.Sprintf("pokemon_%s", pokemonName)

	// Find PokeAPI data in cache, if not found made a new request
	cacheData, found := s.cache.Get(pokeDataKey)
	if found {
		fmt.Println("Pokemon em cache")
		return cacheData.(myTypes.AnyMap), nil
	}
	fmt.Println("Pokemon fora cache")

	pokeJson, err := searchPokemonInPokeApi(pokemonName)
	if err != nil {
		return nil, err
	}

	// Save PokeAPI data in cache permanently
	s.cache.Set(pokeDataKey, pokeJson, cache.NoExpiration)

	return pokeJson, nil
}

func (s *PokemonService) obtainTypeData(types []interface {}) ([]string, *customErrors.InternalError) {

	typesIds := make([]string, 0)
	for _, typeData := range types {
		typeName := (typeData.(myTypes.AnyMap)["type"]).(string)
		typeReferenceId := (strings.Split((typeData.(myTypes.AnyMap)["url"].(string)), "/")[6])

		typeKey := fmt.Sprintf("type_%s", typeName)
		// Find PokeAPI data in cache, if not found made a new request
		cacheData, found := s.cache.Get(typeKey)
		if found {
			typesIds = append(typesIds, cacheData.(string))
		} else {
			typeData, err := s.typesRepository.FindByName(typeName)
			if err != nil {
				return nil, customErrors.NewInternalError("a failure occurred when try to obtain pokémon type data", 500, []string{})
			}
			if typeData == nil {
				referenceId, _ := strconv.ParseUint(typeReferenceId, 10, 0)
				createData := myTypes.CreatePokemonTypeInput{
					ReferenceId: referenceId,
					Name:        typeName,
				}
				typeData, err = s.typesRepository.Create(createData)
				if err != nil {
					return nil, customErrors.NewInternalError("a failure occurred when try to create a new pokémon type data", 500, []string{})
				}
				s.cache.Set(typeKey, typeData.Id, cache.NoExpiration)
			}

			typesIds = append(typesIds, typeData.Id)
		}
	}

	return typesIds, nil
}

func (s *PokemonService) NewPokemon(input myTypes.NewPokemonInput) *customErrors.InternalError {

	if inputErr := validateNewPokemonInput(input); inputErr != nil {
		return inputErr
	}

	pokeData, err := s.obtainPokemonData(strings.ToLower(input.Name))
	if err != nil {
		return err
	}

	typeIds, err := s.obtainTypeData(pokeData["types"].([]interface {}))
	if err != nil {
		return err
	}

	pokemonData := myTypes.CreatePokemonInput{
		ReferenceId: pokeData["id"].(uint),
		TierId:      input.TierId,
		Name:        input.Name,
		Weight:      pokeData["weight"].(uint),
		Height:      pokeData["height"].(uint),
		Experience:  pokeData["base_experience"].(uint),
		ImageUrl:    (((pokeData["sprites"].(myTypes.AnyMap))["other"].(myTypes.AnyMap))["official-artwork"].(myTypes.AnyMap))["front_default"].(string),
	}

	data := myTypes.RegistryPokemonInput{
		CreatePokemonInput: pokemonData,
		Price:              input.Price,
		InitialStock:       input.InitialStock,
		Types:              typeIds,
	}

	_, repositoryErr := s.repository.Registry(data)
	if repositoryErr != nil {
		return customErrors.NewInternalError("a failure occurred when try to registry a new pokemon", 500, []string{})
	}

	return nil
}
