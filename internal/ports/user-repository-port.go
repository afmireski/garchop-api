package ports

import (
	"time"

	"github.com/afmireski/garchop-api/internal/entities"
)

type UserRepositoryPort interface {
	Create(input CreateUserInput) string

	FindById(id string) (interface{}, error)

	Update(id string, input interface{}) (interface{}, error)

	Delete(id string) error
}

type CreateUserInput struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	BirthDate time.Time `json:"birth_date"`
	Role entities.UserRoleEnum `json:"role"`
}