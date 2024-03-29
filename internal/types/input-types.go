package types

import "time"

type NewUserInput struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	BirthDate time.Time `json:"birth_date"`
}

type LoginInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}