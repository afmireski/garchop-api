package ports

import (
	"time"

	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type UserRepositoryPort interface {
	Create(input CreateUserInput) (string, error)

	FindById(id string) (myTypes.Any, error)

	Update(id string, input myTypes.AnyMap, where myTypes.Where) (myTypes.Any, error)

	Delete(id string) error
}

type CreateUserInput struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	PlainPassword string `json:"plain_password"`
	BirthDate time.Time `json:"birth_date"`
}