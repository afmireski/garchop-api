package ports

import (
	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type PokemonRepositoryPort interface {
	Create(input myTypes.CreatePokemonInput) (string, error)

	Registry(input myTypes.RegistryPokemonInput) (string, error)

	FindById(id string, where myTypes.Where) (*models.PokemonModel, error)

	FindAll(where myTypes.Where) ([]models.PokemonModel, error)

	Update(id string, input myTypes.AnyMap, where myTypes.Where) (myTypes.Any, error)
}
