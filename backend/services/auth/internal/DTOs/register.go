package dtos

type RegisterDTO struct {
	Username     string `json:"username"`
	CompleteName string `json:"complete_name"`
	Cellphone    string `json:"cellphone"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}