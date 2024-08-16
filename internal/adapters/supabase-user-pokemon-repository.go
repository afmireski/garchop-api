package adapters

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/nedpals/supabase-go"

	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type SupabaseUserPokemonRepository struct {
	client *supabase.Client
}

func NewSupabaseUserPokemonRepository(client *supabase.Client) *SupabaseUserPokemonRepository {
	return &SupabaseUserPokemonRepository{
		client: client,
	}
}

func (r *SupabaseUserPokemonRepository) serializeToModel(supabaseData myTypes.AnyMap) (*models.UserPokemonModel, error) {
	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		return nil, err
	}

	var modelData models.UserPokemonModel
	err = json.Unmarshal(jsonData, &modelData)
	if err != nil {
		return nil, err
	}

	return &modelData, nil
}

func (r *SupabaseUserPokemonRepository) serializeManyToModel(supabaseData []myTypes.AnyMap) ([]models.UserPokemonModel, error) {
	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		return nil, err
	}

	var modelData []models.UserPokemonModel
	err = json.Unmarshal(jsonData, &modelData)
	if err != nil {
		return nil, err
	}

	return modelData, nil
}

func (r *SupabaseUserPokemonRepository) Upsert(input myTypes.UserPokemonData) (*models.UserPokemonModel, error) {
	var findData myTypes.AnyMap

	findErr := r.client.DB.From("user_pokemons").Select("*").Single().Eq("user_id", input.UserId).Eq("pokemon_id", input.PokemonId).Execute(&findData)
	
	var supabaseData []myTypes.AnyMap
	if findErr != nil {
		if strings.Contains(findErr.Error(), "PGRST116") { // resource not found

			createErr := r.client.DB.From("user_pokemons").Insert(input).Execute(&supabaseData)
			if createErr != nil {
				return nil, createErr
			}

			return r.serializeToModel(supabaseData[0])
		}
		return nil, findErr
	}

	currentQuantity := uint(findData["quantity"].(float64))
	updateData := myTypes.UpdateUserPokemonData{
		Quantity: currentQuantity + input.Quantity,
	}
	fmt.Println(updateData)
	updateErr := r.client.DB.From("user_pokemons").Update(updateData).Eq("user_id", input.UserId).Eq("pokemon_id", input.PokemonId).Execute(&findData)

	if updateErr != nil {
		return nil, updateErr
	}

	return r.serializeToModel(supabaseData[0])
}

func (r *SupabaseUserPokemonRepository) FindById(id string, where myTypes.Where) (*models.UserPokemonModel, error) {
	panic("implement me")
}

func (r *SupabaseUserPokemonRepository) FindAll(where myTypes.Where) ([]models.UserPokemonModel, error) {
	panic("implement me")
}
