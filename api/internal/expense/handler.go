package expense

import (
	"fmt"
	"net/http"
	"strconv"

	dtos "github.com/GeuberLucas/Gofre/api/pkg/DTOs"
	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
	"github.com/GeuberLucas/Gofre/api/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type IHandlerExpense interface {
	AddExpenseHandler(c *fiber.Ctx) error
	GetExpenseHandler(c *fiber.Ctx) error
	GetByIdExpenseHandler(c *fiber.Ctx) error
	UpdateExpenseHandler(c *fiber.Ctx) error
	DeleteExpenseHandler(c *fiber.Ctx) error
	UpdateIsPaidHandler(c *fiber.Ctx) error
}
type HandlerExpenseService struct {
	service IExpenseService
}

func NewHandlerService(service IExpenseService) IHandlerExpense {
	return &HandlerExpenseService{
		service: service,
	}
}
func (h *HandlerExpenseService) AddExpenseHandler(c *fiber.Ctx) error {

	userIdToken := c.Get("user_id")
	userIdInt, err := strconv.ParseInt(userIdToken, 10, 64)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)

	}
	var expenseDto dtos.ExpenseDto
	if err = c.BodyParser(&expenseDto); err != nil {
		return checkErroType(c, err, helpers.VALIDATION)

	}

	expenseDto.UserId = userIdInt
	stringTypeError, err := h.service.AddExpense(expenseDto)
	if err != nil {
		return checkErroType(c, err, stringTypeError)

	}
	var dataReturn interface{}
	return response.JSONResponse(c, http.StatusOK, dataReturn)
}

// GetExpenseHandler implements IHandlerExpense.
func (h *HandlerExpenseService) GetExpenseHandler(c *fiber.Ctx) error {
	panic("unimplemented")
}
func (h *HandlerExpenseService) GetByIdExpenseHandler(c *fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Get("idTransaction"), 10, 64)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)

	}
	serviceresult, typeError, err := h.service.GetByIdExpense(uint(id))
	if err != nil {
		return checkErroType(c, err, typeError)

	}

	return response.JSONResponse(c, http.StatusOK, serviceresult)
}

func (h *HandlerExpenseService) GetByIdUserExpenseHandler(c *fiber.Ctx) error {

	userIdToken := c.Get("user_id")
	userIdInt, err := strconv.ParseUint(userIdToken, 10, 32)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)

	}

	serviceresult, typeError, err := h.service.GetAllExpense(uint(userIdInt))
	if err != nil {
		return checkErroType(c, err, typeError)
	}

	return response.JSONResponse(c, http.StatusOK, serviceresult)
}

func (h *HandlerExpenseService) UpdateExpenseHandler(c *fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Get("idTransaction"), 10, 64)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)

	}
	userIdToken := c.Get("user_id")
	userIdInt, err := strconv.ParseInt(userIdToken, 10, 64)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)

	}

	if err != nil {
		return response.ErrorResponse(c, http.StatusUnprocessableEntity, err)

	}
	var expenseDto dtos.ExpenseDto
	if err = c.BodyParser(&expenseDto); err != nil {
		return checkErroType(c, err, helpers.VALIDATION)

	}
	expenseDto.UserId = userIdInt
	typeError, err := h.service.UpdateExpense(uint(id), expenseDto)
	if err != nil {
		return checkErroType(c, err, typeError)

	}

	return response.JSONResponse(c, http.StatusOK, nil)
}

func (h *HandlerExpenseService) DeleteExpenseHandler(c *fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Get("idTransaction"), 10, 64)
	if err != nil {
		return checkErroType(c, err, helpers.VALIDATION)

	}
	userIdToken := c.Get("user_id")
	userIdInt, err := strconv.ParseUint(userIdToken, 10, 32)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)

	}
	typeError, err := h.service.DeleteExpense(uint(id), uint(userIdInt))
	if err != nil {
		return checkErroType(c, err, typeError)

	}

	return response.JSONResponse(c, http.StatusOK, nil)
}

func (h *HandlerExpenseService) UpdateIsPaidHandler(c *fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Get("idTransaction"), 10, 64)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err)

	}

	if err != nil {
		return response.ErrorResponse(c, http.StatusUnprocessableEntity, err)

	}
	var expenseDto struct {
		IsPaid bool `json:"isPaid"`
	}
	if err = c.BodyParser(&expenseDto); err != nil {
		return checkErroType(c, err, helpers.VALIDATION)

	}
	fmt.Println(expenseDto)
	typeError, err := h.service.UpdateIsPaidExpense(uint(id), expenseDto.IsPaid)
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
