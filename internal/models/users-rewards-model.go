package models

import "time"

type UserRewardModel struct {
	UserId      string       `json:"user_id"`
	RewardId    string       `json:"reward_id"`
	ClaimedAt   time.Time    `json:"claimed_at"`
	RewardModel *RewardModel `json:"reward"`
}
