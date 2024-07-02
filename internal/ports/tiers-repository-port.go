package ports

import (
	"github.com/afmireski/garchop-api/internal/models"

	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type TiersRepositoryPort interface {
	FindAll(where myTypes.Where) ([]models.TierModel, error)

	FindById(id int, where myTypes.Where) (*models.TierModel, error)

	FindWhereUnique(where myTypes.Where) (*models.TierModel, error)
}
