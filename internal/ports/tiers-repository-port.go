package ports

import (
	"github.com/afmireski/garchop-api/internal/models"

	customErrors "github.com/afmireski/garchop-api/internal/errors"
)

type TiersRepositoryPort interface {
	FindAll() ([]models.TierModel, *customErrors.InternalError)

	FindById(id string) (*models.TierModel, *customErrors.InternalError)
}