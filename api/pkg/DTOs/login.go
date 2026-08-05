package dtos

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResultDto struct {
	Token string `json:"token"`
}