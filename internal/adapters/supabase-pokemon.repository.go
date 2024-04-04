package adapters

import (
	supabase "github.com/nedpals/supabase-go"

	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type SupabasePokemonRepository struct {
	client *supabase.Client
}

func NewSupabasePokemonRepository(client *supabase.Client) *SupabasePokemonRepository {
	return &SupabasePokemonRepository{
		client: client,
	}
}

func (r *SupabasePokemonRepository) Create(input myTypes.CreatePokemonInput) (string, error) {
	var supabaseData []map[string]string

	err := r.client.DB.From("pokemons").Insert(input).Execute(&supabaseData); if err != nil {
		return "", err
	}

	return supabaseData[0]["id"], nil
}

type createPriceInput struct {
	pokemonId string `json:"pokemon_id"`
	price int `json:"price"`
}
type createStockInput struct {
	pokemonId string `json:"pokemon_id"`
	quantity int `json:"quantity"`
}
type createPokemonTypeInput struct {
	pokemonId string `json:"pokemon_id"`
	typeId string `json:"type_id"`
}
func (r *SupabasePokemonRepository) Registry(input myTypes.RegistryPokemonInput) (string, error) {
	pokemonId, err := r.Create(input.CreatePokemonInput); if err != nil {
		return "", err
	}

	var supabaseData []map[string]string
	err = r.client.DB.From("prices").Insert(createPriceInput{
		pokemonId: pokemonId,
		price: input.Price,
	}).Execute(&supabaseData); if err != nil {
		return "", err
	}

	err = r.client.DB.From("stocks").Insert(createStockInput{
		pokemonId: pokemonId,
		quantity: input.InitialStock,
	}).Execute(&supabaseData); if err != nil {
		return "", err
	}

	for _, typeId := range input.Types {
		err = r.client.DB.From("pokemon_types").Insert(createPokemonTypeInput{
			pokemonId: pokemonId,
			typeId: typeId,
		}).Execute(&supabaseData); if err != nil {
			return "", err
		}; if err != nil {
			return "", err
		}
	}

	return pokemonId, nil
}

func (r *SupabasePokemonRepository) FindById(id string) (*myTypes.Any, error) {
	panic("method not implemented")
}

func (r *SupabasePokemonRepository) FindAll(where myTypes.Where) ([]myTypes.Any, error) {
	panic("method not implemented")
}

func (r *SupabasePokemonRepository) Update(id string, input myTypes.AnyMap, where myTypes.Where) (myTypes.Any, error) {
	panic("method not implemented")
}




