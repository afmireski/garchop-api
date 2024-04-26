package entities

import "github.com/afmireski/garchop-api/internal/models"

type Item struct {
	Id         string   `json:"id"`
	CartId     string   `json:"cart_id"`
	PokemonId  string   `json:"pokemon_id"`
	PurchaseId *string  `json:"purchase_id",omitempty`
	Quantity   uint     `json:"quantity"`
	Price      uint     `json:"price"`
	Total      uint     `json:"total"`
	Pokemon    *Pokemon `json:"pokemon",omitempty`
}

func BuildItemFromModel(model *models.ItemModel) *Item {
	if model == nil {
		return nil
	}
	return &Item{
		Id:         model.Id,
		CartId:     model.CartId,
		PokemonId:  model.PokemonId,
		PurchaseId: model.PurchaseId,
		Quantity:   model.Quantity,
		Price:      model.Price,
		Total:      model.Total,
		Pokemon:    BuildPokemonFromModel(model.Pokemon),
	}
}
