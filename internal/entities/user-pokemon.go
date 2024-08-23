package entities

import "github.com/afmireski/garchop-api/internal/models"

type UserPokemon struct {
	Pokemon
	UserId    string        `json:"user_id"`
	Quantity  uint          `json:"quantity"`
}

func BuildUserPokemonFromModel(model models.UserPokemonModel) UserPokemon {
	var types []PokemonType
	for _, pokemonType := range model.Pokemon.Types {
		t := PokemonType{Id: pokemonType.Types.Id, Name: pokemonType.Types.Name, ReferenceId: pokemonType.Types.ReferenceId}
		types = append(types, t)
	}

	tier := Tier{
		Id:                model.Pokemon.Tier.Id,
		Name:              model.Pokemon.Tier.Name,
		MinimalExperience: model.Pokemon.Tier.MinimalExperience,
		LimitExperience:   model.Pokemon.Tier.LimitExperience,
	}

	return UserPokemon{
		Pokemon: Pokemon{
			Id:          model.Pokemon.Id,
			ReferenceId: model.Pokemon.ReferenceId,
			Name:        model.Pokemon.Name,
			Weight:      model.Pokemon.Weight,
			Height:      model.Pokemon.Height,
			ImageUrl:    model.Pokemon.ImageUrl,
			Experience:  model.Pokemon.Experience,
			Types:       types,
			Tier:        tier,
		},
		UserId: model.UserId,
		Quantity: model.Quantity,
	}
}

func BuilManyUserPokemonFromModel(models []models.UserPokemonModel) []UserPokemon {
	var userPokemons []UserPokemon
	for _, model := range models {
		userPokemons = append(userPokemons, BuildUserPokemonFromModel(model))
	}

	return userPokemons
}
