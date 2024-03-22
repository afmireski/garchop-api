package ports

import (
	"time"

	"github.com/afmireski/garchop-api/internal/entities"
)

type UserRepositoryPort interface {
	Create(input CreateUserInput) error

	FindById(id string) (interface{}, error)

	Update(id string, input interface{}) (interface{}, error)

	Delete(id string) error
}

type CreateUserInput struct {
	Name string
	Email string
	Phone string
	BirthDate time.Time
	Role entities.UserRoleEnum
}