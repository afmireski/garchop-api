package types

type CreatePurchaseInput struct {
	UserId string `json:"user_id"`
	Cart   string `json:"cart_id"`
}
