package adapters

import (
	"encoding/json"

	supabase "github.com/nedpals/supabase-go"

	"github.com/afmireski/garchop-api/internal/models"
	myTypes "github.com/afmireski/garchop-api/internal/types"
)


type SupabasePurchaseRepository struct {
	supabaseClient *supabase.Client
}

func NewSupabasePurchaseRepository(supabaseClient *supabase.Client) *SupabasePurchaseRepository {
	return &SupabasePurchaseRepository{
		supabaseClient: supabaseClient,
	}
}

func (r *SupabasePurchaseRepository) serializeToModel(supabaseData myTypes.AnyMap) (*models.PurchaseModel, error) {
	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		return nil, err
	}

	var modelData models.PurchaseModel
	err = json.Unmarshal(jsonData, &modelData)
	if err != nil {
		return nil, err
	}

	return &modelData, nil
}

func (r *SupabasePurchaseRepository) serializeManyToModel(supabaseData []myTypes.AnyMap) ([]models.PurchaseModel, error) {
	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		return nil, err
	}

	var modelData []models.PurchaseModel
	err = json.Unmarshal(jsonData, &modelData)
	if err != nil {
		return nil, err
	}

	return modelData, nil
}

func (r *SupabasePurchaseRepository) Create(input myTypes.CreatePurchaseInput) (string, error) {

	var supabaseData []myTypes.AnyMap

	err := r.supabaseClient.DB.From("purchases").Insert(input).Execute(&supabaseData); if err != nil {
		return "", err
	}

	return supabaseData[0]["id"].(string), nil
}

func (r *SupabasePurchaseRepository) Delete(id string) error {
	panic("implement me")
}

func (r *SupabasePurchaseRepository) FindById(id string, where myTypes.Where) (*models.PurchaseModel, error) {
	panic("implement me")
}

func (r *SupabasePurchaseRepository) FindAll(where myTypes.Where) ([]models.PurchaseModel, error) {
	var supabaseData []myTypes.AnyMap

	query := r.supabaseClient.DB.From("purchases").Select("*", "items(*)")

	if len(where) > 0 {
		for column, filter := range where {
			for operator, criteria := range filter {
				query.Filter(column, operator, criteria)
			}
		}
	}

	err := query.Execute(&supabaseData)
	if err != nil {
		return nil, err
	}

	return r.serializeManyToModel(supabaseData)
}
