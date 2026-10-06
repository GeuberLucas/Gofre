package profile

import (
	"net/http"

	dtos "github.com/GeuberLucas/Gofre/api/pkg/DTOs"
	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
	"github.com/GeuberLucas/Gofre/api/pkg/response"
	"github.com/GeuberLucas/Gofre/api/pkg/types" // Import do package de types necessário
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
	userId := c.Locals("user_id").(uint)

	var req dtos.CreateProfileDTO // Usa o novo DTO
	if err := c.Bind().JSON(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)
	}

	model := Profile{
		UserId:         userId,
		CompleteName:   req.CompleteName,
		CellPhone:      req.CellPhone,
		InitialBalance: types.FloatToMoney(req.InitialBalance),
	}

	typeError, err := p.service.AddProfileService(model)
	if err != nil {
		return checkErroType(c, err, typeError)
	}

	return response.JSONResponse(c, http.StatusCreated, fiber.Map{"message": "Profile created successfully"})
}

func (p *ProfileHandler) GetProfileHandler(c fiber.Ctx) error {
	userId := c.Locals("user_id").(uint)

	serviceresult, typeError, err := p.service.GetProfileService(userId)
	if err != nil {
		return checkErroType(c, err, typeError)
	}

	// Usa o Response DTO
	responseObj := dtos.ProfileResponseDTO{
		ID:             serviceresult.ID,
		UserId:         serviceresult.UserId,
		CompleteName:   serviceresult.CompleteName,
		CellPhone:      serviceresult.CellPhone,
		InitialBalance: serviceresult.InitialBalance.ToFloat(),
	}

	return response.JSONResponse(c, http.StatusOK, responseObj)
}

// UpdateProfileHandler implements [IProfileHandler].
func (p *ProfileHandler) UpdateProfileHandler(c fiber.Ctx) error {
	userId := c.Locals("user_id").(uint)

	existingProfile, typeError, err := p.service.GetProfileService(userId)
	if err != nil {
		return checkErroType(c, err, typeError)
	}

	var patchData dtos.UpdateProfileDTO // Usa o novo DTO de atualização
	if err := c.Bind().JSON(&patchData); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)
	}

	if patchData.CompleteName != nil {
		existingProfile.CompleteName = *patchData.CompleteName
	}
	if patchData.CellPhone != nil {
		existingProfile.CellPhone = *patchData.CellPhone
	}
	if patchData.InitialBalance != nil {
		existingProfile.InitialBalance = types.FloatToMoney(*patchData.InitialBalance)
	}

	typeError, err = p.service.UpdateProfileService(existingProfile)
	if err != nil {
		return checkErroType(c, err, typeError)
	}

	// Usa o Response DTO para devolver os dados atualizados
	updatedResponse := dtos.ProfileResponseDTO{
		ID:             existingProfile.ID,
		UserId:         existingProfile.UserId,
		CompleteName:   existingProfile.CompleteName,
		CellPhone:      existingProfile.CellPhone,
		InitialBalance: existingProfile.InitialBalance.ToFloat(),
	}

	return response.JSONResponse(c, http.StatusOK, fiber.Map{
		"message": "Profile patched successfully",
		"data":    updatedResponse,
	})
}

// DeleteProfileHandler implements [IProfileHandler].
func (p *ProfileHandler) DeleteProfileHandler(c fiber.Ctx) error {
	userId := c.Locals("user_id").(uint)

	typeError, err := p.service.DeleteProfileService(userId)
	if err != nil {
		return checkErroType(c, err, typeError)
	}

	return response.JSONResponse(c, http.StatusOK, fiber.Map{"message": "Profile and account deleted successfully"})
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
	case helpers.NOT_FOUND:
		return response.ErrorResponse(c, http.StatusNotFound, err)
	default:
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
}
