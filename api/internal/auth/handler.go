package auth

import (
	"net/http"

	"github.com/GeuberLucas/Gofre/api/internal/auth/security"
	dtos "github.com/GeuberLucas/Gofre/api/pkg/DTOs"
	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
	"github.com/GeuberLucas/Gofre/api/pkg/response"
	"github.com/gofiber/fiber/v3"
)

type IHandlerAuth interface {
	LoginHandler(c fiber.Ctx) error
	RegisterHandler(c fiber.Ctx) error
	ForgotPasswordHandler(c fiber.Ctx) error
	ResetPasswordHandler(c fiber.Ctx) error
	IsAuthenticatedMiddleware(c fiber.Ctx) error
	ProfileHandler(c fiber.Ctx) error
}

type HandlerAuth struct {
	service IAuthService
}

func NewHandlerService(svc IAuthService) IHandlerAuth {
	return &HandlerAuth{
		service: svc,
	}
}
func (h *HandlerAuth) LoginHandler(c fiber.Ctx) error {

	var loginDTO dtos.LoginDTO
	if erro := c.Bind().Body(&loginDTO); erro != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, erro)
	}
	service := h.service
	serviceresult, errorType, erro := service.Login(loginDTO)
	if erro != nil {
		switch errorType {
		case helpers.VALIDATION:
			return response.ErrorResponse(c, http.StatusBadRequest, erro)
		case helpers.INTERNAL:
			return response.ErrorResponse(c, http.StatusInternalServerError, erro)
		default:
			return response.ErrorResponse(c, http.StatusBadRequest, erro)
		}
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt-token",
		Value:    serviceresult.Token,
		Path:     "/",
		Secure:   true,
		HTTPOnly: true,
	})
	return response.JSONResponse(c, http.StatusOK, nil)

}

func (h *HandlerAuth) RegisterHandler(c fiber.Ctx) error {

	var registerDTO dtos.RegisterDTO
	if erro := c.Bind().Body(&registerDTO); erro != nil {

		return response.ErrorResponse(c, http.StatusBadRequest, erro)
	}
	service := h.service
	serviceresult, errorType, erro := service.Register(registerDTO)
	if erro != nil {
		switch errorType {
		case helpers.VALIDATION:
			return response.ErrorResponse(c, http.StatusBadRequest, erro)
		case helpers.INTERNAL:
			return response.ErrorResponse(c, http.StatusInternalServerError, erro)
		default:
			return response.ErrorResponse(c, http.StatusBadRequest, erro)
		}
	}

	return response.JSONResponse(c, http.StatusOK, serviceresult)
}

func (h *HandlerAuth) IsAuthenticatedMiddleware(c fiber.Ctx) error {
	if err := security.ValidateToken(c); err != nil {
		return response.ErrorResponse(c, http.StatusUnauthorized, err)

	}
	userId, err := security.ExtractUserId(c)
	if err != nil {
		return response.ErrorResponse(c, http.StatusUnauthorized, err)

	}
	c.Locals("user_id", userId)
	return c.Next()
}
func (h *HandlerAuth) ProfileHandler(c fiber.Ctx) error {

	userId := c.Locals("user_id").(uint)
	service := h.service
	serviceresult, errorType, erro := service.Profile(userId)
	if erro != nil {
		switch errorType {
		case helpers.VALIDATION:
			return response.ErrorResponse(c, http.StatusBadRequest, erro)
		case helpers.INTERNAL:
			return response.ErrorResponse(c, http.StatusInternalServerError, erro)
		default:
			return response.ErrorResponse(c, http.StatusBadRequest, erro)
		}
	}

	return response.JSONResponse(c, http.StatusOK, serviceresult)
}

func (h *HandlerAuth) ForgotPasswordHandler(c fiber.Ctx) error {

	var forgotPasswordDTO dtos.ForgotPasswordDTO
	if erro := c.Bind().Body(&forgotPasswordDTO); erro != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, erro)

	}
	service := h.service
	_, erro := service.ForgotPassword(forgotPasswordDTO.Email)
	if erro != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, erro)
	}
	return response.JSONResponse(c, http.StatusOK, map[string]string{"message": "If the email exists, a reset link has been sent."})
}

func (h *HandlerAuth) ResetPasswordHandler(c fiber.Ctx) error {

	HashEncoded := c.Get("HashEncoded")
	if HashEncoded == "" {
		return response.ErrorResponse(c, http.StatusBadRequest, nil)

	}
	service := h.service

	var resetPasswordDTO dtos.ResetPasswordDTO
	if erro := c.Bind().Body(&resetPasswordDTO); erro != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, erro)
	}
	_, erro := service.ResetPassword(HashEncoded, resetPasswordDTO.NewPassword)
	if erro != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, erro)
	}
	return response.JSONResponse(c, http.StatusOK, map[string]string{"message": "Password has been reset successfully."})

}
