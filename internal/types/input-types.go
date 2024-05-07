package types

import "time"

type NewUserInput struct {
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	Phone           string     `json:"phone"`
	Password        string     `json:"password"`
	ConfirmPassword string     `json:"confirm_password"`
	BirthDate       *time.Time `json:"birth_date"`
}

type UpdateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdatePokemonInput struct {
	Price  *int `json:"price"`
	Stock  *int `json:"stock"`
	TierId *int `json:"tier_id"`
}

type AddItemToCartInput struct {
	UserId    string `json:"user_id"`
	PokemonId string `json:"pokemon_id"`
	PriceId   string `json:"price_id"`
	Quantity  uint   `json:"quantity"`
}
