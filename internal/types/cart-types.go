package types

import "time"

type CreateCartInput struct {
	UserId    string    `json:"user_id"`
	IsActive  bool      `json:"is_active"`
	ExpiresAt time.Time `json:"expires_at"`
}
