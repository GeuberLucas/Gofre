package dtos

type ForgotPasswordDTO struct {
	Email string `json:"email"`
}

type ResetPasswordDTO struct {
	NewPassword string `json:"new_password"`
}