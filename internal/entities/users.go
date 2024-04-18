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
	Id        string
	Name      string
	Email     string
	Phone     string
	BirthDate time.Time
	Role      UserRoleEnum
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

func BuildManyUserFromModel(data []models.UserModel) []User {
	var result []User

	for _, val := range data {
		result = append(result, *NewUser(val.Id, val.Name, val.Email, val.Phone, val.BirthDate, string(val.Role)))
	}

	return result
}
