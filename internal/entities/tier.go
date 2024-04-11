package entities

type Tier struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	MinimalExperience uint `json:"minimal_experience"`
	LimitExperience uint `json:"limit_experience"`
}
