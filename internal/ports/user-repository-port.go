package ports

import (
	"time"
)

type UserRepositoryPort interface {
	Create(input CreateUserInput) (string, error)

	FindById(id string) (interface{}, error)

	Update(id string, input interface{}) (interface{}, error)

	Delete(id string) error
}

type CreateUserInput struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	BirthDate time.Time `json:"birth_date"`
}