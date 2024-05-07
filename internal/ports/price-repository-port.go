package ports

import "github.com/afmireski/garchop-api/internal/models"

type PriceRepositoryPort interface {
	FindCurrentPrice(pokemonId string) (*models.PriceModel, error)
}