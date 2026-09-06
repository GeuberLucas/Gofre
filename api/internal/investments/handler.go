package investments

import (
	"net/http"
	"strconv"

	dtos "github.com/GeuberLucas/Gofre/api/pkg/DTOs"
	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
	"github.com/GeuberLucas/Gofre/api/pkg/response"
	"github.com/gofiber/fiber/v3"
)

type IHandlerService interface {
	AddInvestmentHandler(c fiber.Ctx) error
	GetInvestmentHandler(c fiber.Ctx) error
	GetByIdInvestmentHandler(c fiber.Ctx) error
	UpdateInvestmentHandler(c fiber.Ctx) error
	DeleteInvestmentHandler(c fiber.Ctx) error
	UpdateIsDoneInvestmentHandler(c fiber.Ctx) error
}

type HandlerService struct {
	portfolioService IPortfolioService
}

func NewHandlerService(service IPortfolioService) IHandlerService {
	return &HandlerService{
		portfolioService: service,
	}
}

func (hd *HandlerService) AddInvestmentHandler(c fiber.Ctx) error {

	var dto dtos.Portfolio

	if err := c.Bind().Body(&dto); err != nil {
		return checkErroType(c, helpers.INTERNAL, err)

	}
	userId := c.Locals("user_id").(uint)
	dto.UserID = userId
	typeError, err := hd.portfolioService.Add(dto)
	if err != nil {
		return checkErroType(c, typeError, err)

	}
	var dataReturn interface{}
	return response.JSONResponse(c, http.StatusOK, dataReturn)

}
func (hd *HandlerService) GetInvestmentHandler(c fiber.Ctx) error {

	userId := c.Locals("user_id").(uint)
	serviceresult, typeError, err := hd.portfolioService.GetAll(userId)
	if err != nil {
		if err != nil {
			return checkErroType(c, typeError, err)

		}

	}

	return response.JSONResponse(c, http.StatusOK, serviceresult)
}

func (hd *HandlerService) GetByIdInvestmentHandler(c fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Get("idInvestment"), 10, 64)
	if err != nil {
		return checkErroType(c, helpers.VALIDATION, err)
	}
	serviceresult, typeError, err := hd.portfolioService.GetById(uint(id))
	if err != nil {
		return checkErroType(c, typeError, err)

	}

	return response.JSONResponse(c, http.StatusOK, serviceresult)
}

func (hd *HandlerService) UpdateInvestmentHandler(c fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Get("idInvestment"), 10, 64)
	if err != nil {
		checkErroType(c, helpers.VALIDATION, err)

	}
	var dto dtos.Portfolio

	if err := c.Bind().Body(&dto); err != nil {
		checkErroType(c, helpers.INTERNAL, err)

	}
	userId := c.Locals("user_id").(uint)
	dto.Id = uint(id)
	dto.UserID = userId
	typeError, err := hd.portfolioService.Update(dto)
	if err != nil {
		checkErroType(c, typeError, err)

	}
	var dataReturn interface{}
	return response.JSONResponse(c, http.StatusOK, dataReturn)
}

func (hd *HandlerService) DeleteInvestmentHandler(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Get("idInvestment"), 10, 64)
	if err != nil {
		checkErroType(c, helpers.VALIDATION, err)

	}
	userId := c.Locals("user_id").(uint)
	typeError, err := hd.portfolioService.Delete(uint(id), userId)
	if err != nil {
		checkErroType(c, typeError, err)

	}
	var dataReturn interface{}
	return response.JSONResponse(c, http.StatusOK, dataReturn)
}

func (hd *HandlerService) UpdateIsDoneInvestmentHandler(c fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Get("idInvestment"), 10, 64)
	if err != nil {
		checkErroType(c, helpers.VALIDATION, err)

	}

	var dto struct {
		IsDone bool `json:"isDone"`
	}

	if err := c.Bind().Body(&dto); err != nil {
		checkErroType(c, helpers.INTERNAL, err)

	}

	typeError, err := hd.portfolioService.UpdateIsDone(uint(id), dto.IsDone)
	if err != nil {
		return checkErroType(c, typeError, err)

	}
	var dataReturn interface{}
	return response.JSONResponse(c, http.StatusOK, dataReturn)
}

func checkErroType(c fiber.Ctx, typeError helpers.ErrorType, err error) error {

	switch typeError {
	case helpers.VALIDATION:
		return response.ErrorResponse(c, http.StatusBadRequest, err)

	case helpers.MISSING:
		return response.ErrorResponse(c, http.StatusNotFound, err)

	case helpers.STATE:
		return response.ErrorResponse(c, http.StatusConflict, err)

	default:
		return response.ErrorResponse(c, http.StatusInternalServerError, err)

	}

}
