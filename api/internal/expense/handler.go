package expense

import (
	"fmt"
	"net/http"
	"strconv"

	dtos "github.com/GeuberLucas/Gofre/api/pkg/DTOs"
	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
	"github.com/GeuberLucas/Gofre/api/pkg/response"
	"github.com/gofiber/fiber/v3"
)

type IHandlerExpense interface {
	AddExpenseHandler(c fiber.Ctx) error
	GetExpenseHandler(c fiber.Ctx) error
	GetByIdExpenseHandler(c fiber.Ctx) error
	UpdateExpenseHandler(c fiber.Ctx) error
	DeleteExpenseHandler(c fiber.Ctx) error
	UpdateIsPaidHandler(c fiber.Ctx) error
}
type HandlerExpenseService struct {
	service IExpenseService
}

func NewHandlerService(service IExpenseService) IHandlerExpense {
	return &HandlerExpenseService{
		service: service,
	}
}
func (h *HandlerExpenseService) AddExpenseHandler(c fiber.Ctx) error {

	userId := c.Locals("user_id").(uint)

	var expenseDto dtos.ExpenseDto
	if err := c.Bind().Body(&expenseDto); err != nil {
		return checkErroType(c, err, helpers.VALIDATION)

	}

	expenseDto.UserId = userId
	stringTypeError, err := h.service.AddExpense(expenseDto)
	if err != nil {
		return checkErroType(c, err, stringTypeError)

	}
	var dataReturn interface{}
	return response.JSONResponse(c, http.StatusOK, dataReturn)
}

func (h *HandlerExpenseService) GetByIdExpenseHandler(c fiber.Ctx) error {
	idExpense := c.Params("idExpense")
	id, err := strconv.ParseInt(idExpense, 10, 64)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)

	}
	serviceresult, typeError, err := h.service.GetByIdExpense(uint(id))
	if err != nil {
		return checkErroType(c, err, typeError)

	}

	return response.JSONResponse(c, http.StatusOK, serviceresult)
}

func (h *HandlerExpenseService) GetExpenseHandler(c fiber.Ctx) error {

	userId := c.Locals("user_id").(uint)

	serviceresult, typeError, err := h.service.GetAllExpense(userId)
	if err != nil {
		return checkErroType(c, err, typeError)
	}

	return response.JSONResponse(c, http.StatusOK, serviceresult)
}

func (h *HandlerExpenseService) UpdateExpenseHandler(c fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Params("idExpense"), 10, 64)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)

	}
	userId := c.Locals("user_id").(uint)

	if err != nil {
		return response.ErrorResponse(c, http.StatusUnprocessableEntity, err)

	}
	var expenseDto dtos.ExpenseDto
	if err = c.Bind().Body(&expenseDto); err != nil {
		return checkErroType(c, err, helpers.VALIDATION)

	}
	expenseDto.UserId = userId
	typeError, err := h.service.UpdateExpense(uint(id), expenseDto)
	if err != nil {
		return checkErroType(c, err, typeError)

	}

	return response.JSONResponse(c, http.StatusOK, nil)
}

func (h *HandlerExpenseService) DeleteExpenseHandler(c fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Params("idExpense"), 10, 64)
	if err != nil {
		return checkErroType(c, err, helpers.VALIDATION)

	}
	userId := c.Locals("user_id").(uint)
	typeError, err := h.service.DeleteExpense(uint(id), userId)
	if err != nil {
		return checkErroType(c, err, typeError)

	}

	return response.JSONResponse(c, http.StatusOK, nil)
}

func (h *HandlerExpenseService) UpdateIsPaidHandler(c fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Params("idExpense"), 10, 64)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)

	}

	if err != nil {
		return response.ErrorResponse(c, http.StatusUnprocessableEntity, err)

	}
	var expenseDto struct {
		IsPaid bool `json:"isPaid"`
	}
	if err = c.Bind().Body(&expenseDto); err != nil {
		return checkErroType(c, err, helpers.VALIDATION)

	}
	fmt.Println(expenseDto)
	typeError, err := h.service.UpdateIsPaidExpense(uint(id), expenseDto.IsPaid)
	if err != nil {
		return checkErroType(c, err, typeError)

	}

	return response.JSONResponse(c, http.StatusOK, nil)
}

func checkErroType(c fiber.Ctx, err error, typeError helpers.ErrorType) error {
	switch typeError {
	case helpers.VALIDATION:
		return response.ErrorResponse(c, http.StatusBadRequest, err)

	default:
		return response.ErrorResponse(c, http.StatusInternalServerError, err)

	}

}
