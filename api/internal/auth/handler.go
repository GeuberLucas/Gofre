package auth

import (
	"net/http"
	"strconv"

	"github.com/GeuberLucas/Gofre/api/internal/auth/security"
	dtos "github.com/GeuberLucas/Gofre/api/pkg/DTOs"
	"github.com/GeuberLucas/Gofre/api/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {

	var loginDTO dtos.LoginDTO
	if erro := c.BodyParser(&loginDTO); erro != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, erro)
	}
	service := NewAuthService()
	serviceresult, erro, typeError := service.Login(loginDTO)
	if erro != nil {
		if typeError == "validation" {

			return response.ErrorResponse(c, http.StatusBadRequest, erro)
		}
		if typeError == "Internal" {
			return response.ErrorResponse(c, http.StatusBadRequest, erro)
		}
		if typeError == "Pass" {
			return response.ErrorResponse(c, http.StatusBadRequest, erro)
		}
	}
	return response.JSONResponse(c, http.StatusOK, serviceresult)

}

func RegisterHandler(c *fiber.Ctx) error {

	var registerDTO dtos.RegisterDTO
	if erro := c.BodyParser(&registerDTO); erro != nil {

		return response.ErrorResponse(c, http.StatusBadRequest, erro)
	}
	service := NewAuthService()
	serviceresult, erro, typeError := service.Register(registerDTO)
	if erro != nil {
		switch typeError {
		case "validation":
			return response.ErrorResponse(c, http.StatusBadRequest, erro)
		case "Internal":

			return response.ErrorResponse(c, http.StatusInternalServerError, erro)
		}

	}

	return response.JSONResponse(c, http.StatusOK, serviceresult)
}

func IsAuthenticatedHandler(c *fiber.Ctx) error {
	if err := security.ValidateToken(c); err != nil {
		return response.ErrorResponse(c, http.StatusUnauthorized, err)

	}

	user_Id, err := security.ExtractUserId(c)
	if err != nil {
		return response.ErrorResponse(c, http.StatusUnauthorized, err)

	}
	var userAuthenticated dtos.UserAuthenticatedDto
	userAuthenticated.UserId = uint(user_Id)

	return response.JSONResponse(c, http.StatusOK, userAuthenticated)
}
func ProfileHandler(c *fiber.Ctx) error {

	userId, erro := strconv.ParseInt(c.Get("userId"), 10, 64)
	if erro != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, erro)

	}
	service := NewAuthService()
	serviceresult, erro, typeError := service.Profile(userId)
	if erro != nil {
		if typeError == "Validation" {
			return response.ErrorResponse(c, http.StatusBadRequest, erro)

		}
		if typeError == "Internal" {
			return response.ErrorResponse(c, http.StatusInternalServerError, erro)

		}

	}

	return response.JSONResponse(c, http.StatusOK, serviceresult)
}

func ForgotPasswordHandler(c *fiber.Ctx) error {

	var forgotPasswordDTO dtos.ForgotPasswordDTO
	if erro := c.BodyParser(&forgotPasswordDTO); erro != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, erro)

	}
	service := NewAuthService()
	erro := service.ForgotPassword(forgotPasswordDTO.Email)
	if erro != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, erro)

	}
	return response.JSONResponse(c, http.StatusOK, map[string]string{"message": "If the email exists, a reset link has been sent."})
}

func ResetPasswordHandler(c *fiber.Ctx) error {

	HashEncoded := c.Get("HashEncoded")
	if HashEncoded == "" {
		return response.ErrorResponse(c, http.StatusBadRequest, nil)

	}
	service := NewAuthService()

	var resetPasswordDTO dtos.ResetPasswordDTO
	if erro := c.BodyParser(&resetPasswordDTO); erro != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, erro)

	}
	erro := service.ResetPassword(HashEncoded, resetPasswordDTO.NewPassword)
	if erro != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, erro)

	}
	return response.JSONResponse(c, http.StatusOK, map[string]string{"message": "Password has been reset successfully."})

}
