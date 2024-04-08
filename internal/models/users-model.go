package models

import (
	"time"

	"github.com/afmireski/garchop-api/internal/entities"
)

type UserModel struct {
	Id        string                `json:"id"`
	Name      string                `json:"name"`
	Email     string                `json:"email"`
	Password  string                `json:"password"`
	Phone     string                `json:"phone"`
	BirthDate time.Time             `json:"birth_date"`
	Role      entities.UserRoleEnum `json:"role"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
	DeletedAt time.Time             `json:"deleted_at"`
}

func NewUserModel(id string, name string, email string, phone string, birthDate time.Time, role entities.UserRoleEnum,
	createdAt time.Time, updatedAt time.Time, deletedAt time.Time) *UserModel {
	return &UserModel{
		Id:        id,
		Name:      name,
		Email:     email,
		Phone:     phone,
		BirthDate: birthDate,
		Role:      role,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}
