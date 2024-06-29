package entities

import "github.com/afmireski/garchop-api/internal/models"

type Tier struct {
	Id                uint          `json:"id"`
	PreviousTierId    uint          `json:"previous_tier_id"`
	Name              string        `json:"name"`
	MinimalExperience uint          `json:"minimal_experience"`
	LimitExperience   uint          `json:"limit_experience"`
	PreviousTier      *PreviousTier `json:"previous_tier",omitempty`
}

type PreviousTier struct {
	Id                uint   `json:"id"`
	Name              string `json:"name"`
	MinimalExperience uint   `json:"minimal_experience"`
	LimitExperience   uint   `json:"limit_experience"`
}

func BuildTierFromModel(model *models.TierModel) *Tier {
	if model == nil {
		return nil
	}

	return &Tier{
		Id:                model.Id,
		PreviousTierId:    model.PreviousTierId,
		Name:              model.Name,
		MinimalExperience: model.MinimalExperience,
		LimitExperience:   model.LimitExperience,
		PreviousTier:      BuildPreviousTierFromModel(model.PreviousTier),
	}
}

func BuildPreviousTierFromModel(model *models.TierModel) *PreviousTier {
	if model == nil {
		return nil
	}

	return &PreviousTier{
		Id:                model.Id,
		Name:              model.Name,
		MinimalExperience: model.MinimalExperience,
		LimitExperience:   model.LimitExperience,
	}
}

func BuildTiersFromModels(models []models.TierModel) []Tier {
	var tiers []Tier
	for _, model := range models {
		tier := BuildTierFromModel(&model)
		tiers = append(tiers, *tier)
	}
	return tiers
}
