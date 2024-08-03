package adapters

import (
	"encoding/json"

	"github.com/afmireski/garchop-api/internal/models"
	"github.com/nedpals/supabase-go"

	myTypes "github.com/afmireski/garchop-api/internal/types"
)

type SupabasePaymentMethodsRepository struct {
	client *supabase.Client
}

func NewSupabasePaymentMethodsRepository(client *supabase.Client) *SupabasePaymentMethodsRepository {
	return &SupabasePaymentMethodsRepository{
		client: client,
	}
}

func (r *SupabasePaymentMethodsRepository) serializeManySupabaseDataToModel(supabaseData []myTypes.AnyMap) ([]models.PaymentMethodModel, error) {
	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		return nil, err
	}

	var modelData []models.PaymentMethodModel
	err = json.Unmarshal(jsonData, &modelData)
	if err != nil {
		return nil, err
	}

	return modelData, nil
}

func (r *SupabasePaymentMethodsRepository) FindAll(where myTypes.Where) ([]models.PaymentMethodModel, error) {
	var supabaseData []myTypes.AnyMap

	query := r.client.DB.From("payment_methods").Select("*").Is("deleted_at", "null")

	if len(where) > 0 {
		for column, filter := range where {
			for operator, criteria := range filter {
				query = query.Filter(column, operator, criteria)
			}
		}
	}

	err := query.Execute(&supabaseData)
	if err != nil {
		return nil, err
	}

	return r.serializeManySupabaseDataToModel(supabaseData)
}
