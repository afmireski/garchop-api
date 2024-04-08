package types

type LoginOutput struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UpdateClientOutput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
