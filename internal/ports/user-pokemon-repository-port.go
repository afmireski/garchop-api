package ports

import (
	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type UserPokemonRepositoryPort interface {
	Upsert(input myTypes.UserPokemonId) (*models.UserPokemonModel, error)
	FindById(id string, where myTypes.Where) (*models.UserPokemonModel, error)
	FindAll(where myTypes.Where) ([]models.UserPokemonModel, error)
}
