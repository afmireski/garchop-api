package types

import (
	"encoding/json"
	"time"

	"github.com/afmireski/garchop-api/internal/types/enums"
)

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
	Quantity  uint   `json:"quantity"`
}

type AddItemToCartBody struct {
	PokemonId string `json:"pokemon_id"`
	Quantity  uint   `json:"quantity"`
}

type RemoveItemFromCartInput struct {
	ItemId string `json:"item_id"`
	CartId string `json:"cart_id"`
}

type UpdateItemInCartInput struct {
	ItemId   string `json:"item_id"`
	CartId   string `json:"cart_id"`
	Quantity uint   `json:"quantity"`
}

type UpdateItemInCartBody struct {
	Quantity uint   `json:"quantity"`
}

type NewRewardInput struct {
	TierId             uint                `json:"tier_id"`
	Name               string              `json:"name"`
	Description        string              `json:"description"`
	ExperienceRequired uint                `json:"experience_required"`
	Type               enums.PrizeTypeEnum `json:"prize_type"`
	Prize              json.RawMessage     `json:"prize"`
}
