package adapters

import (
	"encoding/json"
	"strings"
	"time"

	supabase "github.com/nedpals/supabase-go"

	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type SupabaseCartsRepository struct {
	client *supabase.Client
}

func NewSupabaseCartsRepository(client *supabase.Client) *SupabaseCartsRepository {
	return &SupabaseCartsRepository{
		client: client,
	}
}

func (r *SupabaseCartsRepository) serializeToModel(supabaseData myTypes.AnyMap) (*models.CartModel, error) {
	if supabaseData["users"] != nil {
		(supabaseData["users"].(myTypes.AnyMap))["birth_date"], _ = time.Parse("2006-01-02", (supabaseData["users"].(myTypes.AnyMap))["birth_date"].(string))
	}

	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		return nil, err
	}

	var modelData models.CartModel
	err = json.Unmarshal(jsonData, &modelData)
	if err != nil {
		return nil, err
	}

	return &modelData, nil
}

func (r *SupabaseCartsRepository) serializeToItensModel(supabaseData myTypes.Any) (*models.ItemModel, error) {
	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		return nil, err
	}

	var modelData models.ItemModel
	err = json.Unmarshal(jsonData, &modelData)
	if err != nil {
		return nil, err
	}

	return &modelData, nil
}

func (r *SupabaseCartsRepository) Create(input myTypes.CreateCartInput) (*models.CartModel, error) {
	var supabaseData []myTypes.AnyMap

	err := r.client.DB.From("carts").Insert(input).Execute(&supabaseData);

	if err != nil {
		return nil, err
	}

	return r.serializeToModel(supabaseData[0])
}

func (r *SupabaseCartsRepository) FindById(id string, where myTypes.Where) (*models.CartModel, error) {
	var supabaseData myTypes.AnyMap

	query := r.client.DB.From("carts").Select("*", "items(*, pokemons (*, pokemon_types (*, types (*)), tiers (*)))", "users(*, user_stats(*))").Single().Eq("id", id)

	if len(where) > 0 {
		for column, filter := range where {
			for operator, criteria := range filter {
				query = query.Filter(column, operator, criteria)
			}
		}
	} 

	err := query.Execute(&supabaseData); if err != nil {
		if strings.Contains(err.Error(), "PGRST116") { // resource not found
			return nil, nil
		}

		return nil, err
	}
	return r.serializeToModel(supabaseData)
}

func (r *SupabaseCartsRepository) FindLastCart(user_id string) (*models.CartModel, error) {
	var supabaseData myTypes.AnyMap

	err := r.client.DB.From("carts").Select("*").OrderBy("created_at", "desc").Limit(1).Single().Eq("user_id", user_id).Is("deleted_at", "null").Eq("is_active", "true").Gt("expires_at", "now()").Execute(&supabaseData)
	if err != nil {
		if strings.Contains(err.Error(), "PGRST116") { // resource not found
			return nil, nil
		}

		return nil, err
	}

	cartData, err := r.serializeToModel(supabaseData); if err != nil {
		return nil, err
	}

	return cartData, nil
}

func (r *SupabaseCartsRepository) Update(id string, input myTypes.AnyMap, where myTypes.Where) (*models.CartModel, error) {
	var supabaseData []myTypes.AnyMap

	query := r.client.DB.From("carts").Update(input).Eq("id", id);

	if len(where) > 0 {
		for column, filter := range where {
			for operator, criteria := range filter {
				query = query.Filter(column, operator, criteria)
			}
		}
	}

	err := query.Execute(&supabaseData); if err != nil {
		return nil, err
	}

	return r.serializeToModel(supabaseData[0])
}

func (r *SupabaseCartsRepository) Delete(id string) error {
	var supabaseData []myTypes.AnyMap

	return r.client.DB.From("carts").Delete().Eq("id", id).Execute(&supabaseData)
}
