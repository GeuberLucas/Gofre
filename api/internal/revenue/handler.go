package revenue

import (
	"net/http"
	"strconv"

	dtos "github.com/GeuberLucas/Gofre/api/pkg/DTOs"
	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
	"github.com/GeuberLucas/Gofre/api/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type IRevenueHandler interface {
	AddRevenueHandler(c *fiber.Ctx) error
	GetRevenueHandler(c *fiber.Ctx) error
	GetByIdRevenueHandler(c *fiber.Ctx) error
	UpdateRevenueHandler(c *fiber.Ctx) error
	DeleteRevenueHandler(c *fiber.Ctx) error
	UpdateIsRecievedHandler(c *fiber.Ctx) error
}

type HandlerRevenueService struct {
	service IRevenueService
}

func NewHandlerService(service IRevenueService) IRevenueHandler {
	return &HandlerRevenueService{
		service: service,
	}
}

func (h *HandlerRevenueService) AddRevenueHandler(c *fiber.Ctx) error {
	userIdToken := c.Get("user_id")
	userIdInt, err := strconv.ParseInt(userIdToken, 10, 64)

	var revenueDto dtos.RevenueDto
	if err = c.BodyParser(&revenueDto); err != nil {
		return checkErroType(c, err, helpers.VALIDATION)
	}
	revenueDto.UserId = userIdInt
	typeError, err := h.service.AddRevenue(revenueDto)
	if err != nil {
		return checkErroType(c, err, typeError)
	}
	return response.JSONResponse(c, http.StatusOK, nil)
}
func (h *HandlerRevenueService) GetByIdRevenueHandler(c *fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Get("idTransaction"), 10, 64)
	if err != nil {
		return checkErroType(c, err, helpers.VALIDATION)
	}
	serviceresult, err, typeError := h.service.GetByIdRevenue(id)
	if err != nil {
		return checkErroType(c, err, typeError)

	}

	return response.JSONResponse(c, http.StatusOK, serviceresult)
}

func (h *HandlerRevenueService) GetRevenueHandler(c *fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Get("idTransaction"), 10, 64)
	if err != nil {
		return checkErroType(c, err, helpers.VALIDATION)
	}
	serviceresult, err, typeError := h.service.GetAllRevenue(id)
	if err != nil {
		return checkErroType(c, err, typeError)

	}

	return response.JSONResponse(c, http.StatusOK, serviceresult)
}

func (h *HandlerRevenueService) GetByIdUserRevenueHandler(c *fiber.Ctx) error {
	userIdToken := c.Get("user_id")
	userIdInt, err := strconv.ParseInt(userIdToken, 10, 64)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)

	}
	serviceresult, err, typeError := h.service.GetByIdUserRevenue(userIdInt)
	if err != nil {
		return checkErroType(c, err, typeError)
	}

	return response.JSONResponse(c, http.StatusOK, serviceresult)
}

func (h *HandlerRevenueService) UpdateRevenueHandler(c *fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Get("idTransaction"), 10, 64)
	if err != nil {
		return checkErroType(c, err, helpers.VALIDATION)
	}
	userIdToken := c.Get("user_id")
	userIdInt, err := strconv.ParseInt(userIdToken, 10, 64)

	var revenueDto dtos.RevenueDto
	revenueDto.UserId = userIdInt
	if err = c.BodyParser(&revenueDto); err != nil {
		return checkErroType(c, err, helpers.VALIDATION)
	}
	err, typeError := h.service.UpdateRevenue(id, revenueDto)
	if err != nil {
		return checkErroType(c, err, typeError)

	}

	return response.JSONResponse(c, http.StatusOK, nil)
}

func (h *HandlerRevenueService) DeleteRevenueHandler(c *fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Get("idTransaction"), 10, 64)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)

	}
	userIdToken := c.Get("user_id")
	userIdInt, err := strconv.ParseInt(userIdToken, 10, 64)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)

	}
	err, typeError := h.service.DeleteRevenue(id, userIdInt)
	if err != nil {

		return checkErroType(c, err, typeError)
	}

	return response.JSONResponse(c, http.StatusOK, nil)
}

func (h *HandlerRevenueService) UpdateIsRecievedHandler(c *fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Get("idTransaction"), 10, 64)
	if err != nil {
		return checkErroType(c, err, helpers.VALIDATION)
	}

	var revenueDto struct {
		IsRecieved bool `json:"isRecieved"`
	}
	if err = c.BodyParser(&revenueDto); err != nil {
		return checkErroType(c, err, helpers.VALIDATION)
	}
	err, typeError := h.service.UpdateIsReceivedRevenue(id, revenueDto.IsRecieved)
	if err != nil {
		return checkErroType(c, err, typeError)

	}

	return response.JSONResponse(c, http.StatusOK, nil)
}

func checkErroType(c *fiber.Ctx, err error, typeError helpers.ErrorType) error {
	switch typeError {
	case helpers.VALIDATION:
		return response.ErrorResponse(c, http.StatusBadRequest, err)

	default:
		return response.ErrorResponse(c, http.StatusInternalServerError, err)

	}

}
