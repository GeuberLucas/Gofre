package expense

import (
	"fmt"

	dtos "github.com/GeuberLucas/Gofre/api/pkg/DTOs"
	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
	"github.com/GeuberLucas/Gofre/api/pkg/types"
)

type IExpenseService interface {
	AddExpense(dto dtos.ExpenseDto) (helpers.ErrorType, error)
	GetAllExpense(userId uint) ([]dtos.ExpenseDto, helpers.ErrorType, error)
	GetByIdExpense(id uint) (dtos.ExpenseDto, helpers.ErrorType, error)
	UpdateExpense(id uint, dto dtos.ExpenseDto) (helpers.ErrorType, error)
	DeleteExpense(id uint, userId uint) (helpers.ErrorType, error)
	UpdateIsPaidExpense(id uint, isDone bool) (helpers.ErrorType, error)
}
type ExpenseService struct {
	expenseRepository IExpenseRepository
}

func NewExpenseService(repository IExpenseRepository) IExpenseService {
	return &ExpenseService{
		expenseRepository: repository,
	}
}
func (ts *ExpenseService) AddExpense(dto dtos.ExpenseDto) (helpers.ErrorType, error) {
	model := toModel(dto)
	err := model.Isvalid()
	if err != nil {
		return helpers.VALIDATION, err
	}

	err = ts.expenseRepository.Create(model)
	if err != nil {
		return helpers.INTERNAL, err
	}

	return helpers.NONE, nil
}

func (ts *ExpenseService) GetByIdExpense(id uint) (dtos.ExpenseDto, helpers.ErrorType, error) {
	expenseModel, err := ts.expenseRepository.GetById(id)
	if err != nil {
		return dtos.ExpenseDto{}, helpers.INTERNAL, err
	}
	expenseDto := expenseDtoFromModel(expenseModel)
	return expenseDto, helpers.NONE, nil
}
func (ts *ExpenseService) GetAllExpense(idUser uint) ([]dtos.ExpenseDto, helpers.ErrorType, error) {
	expenseModels, err := ts.expenseRepository.GetAll(idUser)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}
	var expensesDtos []dtos.ExpenseDto
	for _, expenseModel := range expenseModels {
		expenseDto := expenseDtoFromModel(expenseModel)
		expensesDtos = append(expensesDtos, expenseDto)
	}
	return expensesDtos, helpers.NONE, nil
}
func (ts *ExpenseService) UpdateExpense(id uint, dto dtos.ExpenseDto) (helpers.ErrorType, error) {

	model := toModel(dto)
	model.ID = id
	err := model.Isvalid()
	if err != nil {
		return helpers.VALIDATION, err
	}
	err = ts.expenseRepository.Update(model)
	if err != nil {
		return helpers.INTERNAL, err
	}

	return helpers.NONE, nil
}

func (ts *ExpenseService) DeleteExpense(id uint, userId uint) (helpers.ErrorType, error) {

	err := ts.expenseRepository.Delete(id, userId)
	if err != nil {
		return helpers.INTERNAL, err
	}

	return helpers.NONE, nil
}

func (ts *ExpenseService) UpdateIsPaidExpense(id uint, isPaid bool) (helpers.ErrorType, error) {
	oldModel, err := ts.expenseRepository.GetById(id)

	if err != nil {
		return helpers.INTERNAL, err
	}
	model := oldModel
	model.ID = id
	model.IsPaid = isPaid
	fmt.Println(isPaid)
	fmt.Println(model)
	err = model.Isvalid()
	if err != nil {
		return helpers.VALIDATION, err
	}
	err = ts.expenseRepository.Update(model)
	if err != nil {
		return helpers.INTERNAL, err
	}

	return helpers.NONE, nil
}

func expenseDtoFromModel(ex Expense) dtos.ExpenseDto {
	return dtos.ExpenseDto{
		ID:            int64(ex.ID),
		UserId:        int64(ex.UserId),
		Description:   ex.Description,
		Target:        ex.Target,
		Category:      ex.Category,
		Type:          ex.Type,
		PaymentMethod: ex.PaymentMethod,
		PaymentDate:   ex.PaymentDate,
		IsPaid:        ex.IsPaid,
		Amount:        ex.Amount.ToFloat(),
	}
}
func toModel(ex dtos.ExpenseDto) Expense {
	amount := types.FloatToMoney(ex.Amount)
	return Expense{
		ID:            uint(ex.ID),
		UserId:        uint(ex.UserId),
		Description:   ex.Description,
		Target:        ex.Target,
		Category:      ex.Category,
		Type:          ex.Type,
		PaymentMethod: ex.PaymentMethod,
		PaymentDate:   ex.PaymentDate,
		IsPaid:        ex.IsPaid,
		Amount:        amount,
	}
}
