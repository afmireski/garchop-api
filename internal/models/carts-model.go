package models

import "time"

type CartModel struct {
	Id        string      `json:"id"`
	UserId    string      `json:"user_id"`
	IsActive  bool        `json:"is_active"`
	ExpiresIn time.Time   `json:"expires_in"`
	Total     uint        `json:"total"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt *time.Time  `json:"deleted_at,omitempty"`
	Items     []ItemModel `json:"items"`
}
