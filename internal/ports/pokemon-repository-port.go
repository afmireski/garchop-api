package ports

import myTypes "github.com/afmireski/garchop-api/internal/types"

type PokemonRepositoryPort interface {
	Create(input myTypes.CreatePokemonInput) (string, error)

	Registry(input myTypes.RegistryPokemonInput) (string, error)

	FindById(id string) (*myTypes.Any, error)

	FindAll(where myTypes.Where) ([]myTypes.Any, error)

	Update(id string, input myTypes.AnyMap, where myTypes.Where) (myTypes.Any, error)
}
