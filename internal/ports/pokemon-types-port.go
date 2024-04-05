package ports

import (
	"github.com/afmireski/garchop-api/internal/entities"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type PokemonTypesRepositoryPort interface {
	Create(input myTypes.CreatePokemonTypeInput) (*entities.PokemonType, error)

	FindByName(name string) (*entities.PokemonType, error)
}
