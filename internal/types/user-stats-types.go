package types

type CreateUserStatsInput struct {
	UserId string `json:"user_id"`
	TierId uint   `json:"tier_id"`
}
