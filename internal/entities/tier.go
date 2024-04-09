package entities

type Tier struct {
	Id string `json:"id"`
	ReferenceId uint `json:"reference_id"`
	Name string `json:"name"`
	MinimalExperience uint `json:"minimal_experience"`
	LimitExperience uint `json:"limit_experience"`
}
