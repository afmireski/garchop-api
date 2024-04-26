package entities

import (
	"time"

	"github.com/afmireski/garchop-api/internal/models"
)

type Cart struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	IsActive  bool      `json:"is_active"`
	ExpiresIn time.Time `json:"expires_in"`
	Total     uint      `json:"total"`
	Items     []Item    `json:"items"`
}

func BuildCartFromModel(model models.CartModel) *Cart {

	var items []Item
	for _, itemData := range model.Items {
		item := BuildItemFromModel(&itemData);
		items = append(items, *item)
	}

	return &Cart{
		Id:        model.Id,
		UserId:    model.UserId,
		IsActive:  model.IsActive,
		ExpiresIn: model.ExpiresIn,
		Total:     model.Total,
		Items:     items,
	}
}
