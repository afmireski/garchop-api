package ports

import (
	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type UserPokemonRepositoryPort interface {
	Upsert(input myTypes.UserPokemonId) *models.UserPokemonModel
	FindById(id string, where myTypes.Where) (*models.UserPokemonModel, error)
	FindAll(where myTypes.Where) ([]models.UserPokemonModel, error)
	Update(id myTypes.UserPokemonId, where myTypes.Where) error
}
