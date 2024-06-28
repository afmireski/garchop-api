package entities

import (
	"time"

	"github.com/afmireski/garchop-api/internal/models"
)

type UserRoleEnum string

const (
	Client UserRoleEnum = "client"
	Admin  UserRoleEnum = "admin"
)

type User struct {
	Id        string       `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Phone     string       `json:"phone"`
	BirthDate time.Time    `json:"birth_date"`
	Role      UserRoleEnum `json:"role"`
	Status    *UserStats   `json:"status",omitempty`
}

func NewUser(id string, name string, email string, phone string, birthDate time.Time, role string) *User {
	return &User{
		Id:        id,
		Name:      name,
		Email:     email,
		Phone:     phone,
		BirthDate: birthDate,
		Role:      UserRoleEnum(role),
	}
}

func BuildUserFromModel(model models.UserModel) *User {
	return &User{
		Id:        model.Id,
		Name:      model.Name,
		Email:     model.Email,
		Phone:     model.Phone,
		BirthDate: model.BirthDate,
		Role:      UserRoleEnum(model.Role),
		Status:    BuildUserStatsFromModel(*model.Stats),
	}
}

func BuildManyUserFromModel(data []models.UserModel) []User {
	result := []User{}

	for _, val := range data {
		result = append(result, *NewUser(val.Id, val.Name, val.Email, val.Phone, val.BirthDate, string(val.Role)))
	}

	return result
}
