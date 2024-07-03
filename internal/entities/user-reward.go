package entities

type UserReward struct {
	UserId    string  `json:"user_id"`
	RewardId  string  `json:"reward_id"`
	ClaimedAt string  `json:"claimed_at"`
	Reward    *Reward `json:"reward"`
}
