package profile

import (
	"net/http"

	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
	"github.com/GeuberLucas/Gofre/api/pkg/response"
	"github.com/gofiber/fiber/v3"
)

type IProfileHandler interface {
	GetProfileHandler(c fiber.Ctx) error
	AddProfileHandler(c fiber.Ctx) error
	UpdateProfileHandler(c fiber.Ctx) error
	DeleteProfileHandler(c fiber.Ctx) error
}

type ProfileHandler struct {
	service IProfileService
}

// AddProfileHandler implements [IProfileHandler].
func (p *ProfileHandler) AddProfileHandler(c fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteProfileHandler implements [IProfileHandler].
func (p *ProfileHandler) DeleteProfileHandler(c fiber.Ctx) error {
	panic("unimplemented")
}

// GetProfileHandler implements [IProfileHandler].
func (p *ProfileHandler) GetProfileHandler(c fiber.Ctx) error {
	userId := c.Locals("user_id").(uint)

	serviceresult, typeError, err := p.service.GetProfileService(userId)
	if err != nil {
		return checkErroType(c, err, typeError)
	}

	return response.JSONResponse(c, http.StatusOK, serviceresult)
}

// UpdateProfileHandler implements [IProfileHandler].
func (p *ProfileHandler) UpdateProfileHandler(c fiber.Ctx) error {
	panic("unimplemented")
}

func NewHandlerService(service IProfileService) IProfileHandler {
	return &ProfileHandler{
		service: service,
	}
}

func checkErroType(c fiber.Ctx, err error, typeError helpers.ErrorType) error {
	switch typeError {
	case helpers.VALIDATION:
		return response.ErrorResponse(c, http.StatusBadRequest, err)

	default:
		return response.ErrorResponse(c, http.StatusInternalServerError, err)

	}

}
