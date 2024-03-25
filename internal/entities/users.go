package entities

import "time"

type UserRoleEnum string

const (
	Client UserRoleEnum = "client"
	Admin UserRoleEnum = "admin"
)

type User struct {
	Id string
	Name string
	Email string
	Phone string
	BirthDate time.Time
	Role UserRoleEnum
}