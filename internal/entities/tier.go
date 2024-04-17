package entities

import "github.com/afmireski/garchop-api/internal/models"

type Tier struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	MinimalExperience uint `json:"minimal_experience"`
	LimitExperience uint `json:"limit_experience"`
}

func BuildTierFromModel(model models.TierModel) Tier {
	return Tier{
		Id: model.Id,
		Name: model.Name,
		MinimalExperience: model.MinimalExperience,
		LimitExperience: model.LimitExperience,
	}
}

func BuildTiersFromModels(models []models.TierModel) []Tier {
	var tiers []Tier
	for _, model := range models {
		tier := BuildTierFromModel(model)
		tiers = append(tiers, tier)
	}
	return tiers
}
