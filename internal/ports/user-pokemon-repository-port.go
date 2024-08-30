package ports

import (
	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type UserPokemonRepositoryPort interface {
	Upsert(input myTypes.UserPokemonData) (*models.UserPokemonModel, error)
	FindById(userId string, pokemonId string, where myTypes.Where) (*models.UserPokemonModel, error)
	FindAll(where myTypes.Where) ([]models.UserPokemonModel, error)
}
