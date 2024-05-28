package models

import "time"

type PurchaseModel struct {
	Id               string     `json:"id"`
	UserId           string     `json:"user_id"`
	PaymentMethodId  string     `json:"payment_method_id"`
	Total            int        `json:"total"`
	PaymentLimitTime *time.Time `json:"payment_limit_time"`
	IsApproved       bool       `json:"is_approved"`
	CreatedAt        time.Time  `json:"created_at"`
}
