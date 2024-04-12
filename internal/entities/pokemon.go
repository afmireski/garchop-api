package entities

import (
	"github.com/afmireski/garchop-api/internal/models"
)

type Pokemon struct {
	Id          string        `json:"id"`
	ReferenceId uint          `json:"reference_id"`
	Name        string        `json:"name"`
	Weight      uint          `json:"weight"`
	Height      uint          `json:"height"`
	ImageUrl    string        `json:"image_url"`
	Experience  uint          `json:"experience"`
	Types       []PokemonType `json:"types"`
	Tier        Tier          `json:"tier"`
}

type PokemonProduct struct {
	Pokemon
	Price   uint `json:"price"`
	InStock uint `json:"in_stock"`
}

func BuildPokemonProductFromModel(model models.PokemonModel) *PokemonProduct {
	var types []PokemonType
	for _, pokemonType := range model.Types {
		t := PokemonType{Id: pokemonType.Types.Id, Name: pokemonType.Types.Name, ReferenceId: pokemonType.Types.ReferenceId}
		types = append(types, t)
	}

	tier := Tier{
		Id:                model.Tier.Id,
		Name:              model.Tier.Name,
		MinimalExperience: model.Tier.MinimalExperience,
		LimitExperience:   model.Tier.LimitExperience,
	}

	return &PokemonProduct{
		Pokemon: Pokemon{
			Id:          model.Id,
			ReferenceId: model.ReferenceId,
			Name:        model.Name,
			Weight:      model.Weight,
			Height:      model.Height,
			ImageUrl:    model.ImageUrl,
			Experience:  model.Experience,
			Types:       types,
			Tier:        tier,
		},
		Price:   model.Prices[0].Value,
		InStock: model.Stock.Quantity,
	}
}

func BuildManyPokemonProductFromModel(data []models.PokemonModel) []PokemonProduct {
	var result []PokemonProduct

	for _, val := range data {
		var tmp *PokemonProduct
		tmp = BuildPokemonProductFromModel(val)
		result = append(result, *tmp)
	}

	return result
}
