package models

import "time"

type TypeModel struct {
	Id string `json:"id"`
	ReferenceId uint `json:"reference_id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}