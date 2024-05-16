package types

import "time"

type CreatePurchaseInput struct {
	UserId           string     `json:"user_id"`
	CartId           string     `json:"cart_id"`
	PaymentMethodId  *string    `json:"payment_method_id"`
	Total            uint       `json:"total"`
	PaymentLimitDate time.Time `json:"payment_limit_time"`
}

type FinishPurchaseInput struct {
	UserId string `json:"user_id"`
	CartId string `json:"cart_id"`
	PaymentMethodId  *string    `json:"payment_method_id"`
}
