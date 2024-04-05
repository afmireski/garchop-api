package entities

import "time"

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

func NewUser(id string, name string, email string, phone string, birthDate time.Time, role UserRoleEnum) *User {
	return &User{
		Id:        id,
		Name:      name,
		Email:     email,
		Phone:     phone,
		BirthDate: birthDate,
		Role:      role,
	}
}
