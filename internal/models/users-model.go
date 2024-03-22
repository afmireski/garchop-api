package models

import (
	"time"

	"github.com/afmireski/garchop-api/internal/entities"
)

type UserModel struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	BirthDate time.Time `json:"birth_date"`
	Role entities.UserRoleEnum `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}