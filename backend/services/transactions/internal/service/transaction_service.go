package service

import (
	"errors"

	dtos "github.com/GeuberLucas/Gofre/backend/services/transaction/internal/Dtos"
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/models"
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/repository"
)

type TransactionService struct {
	revenueRepository repository.IRevenueRepository
	expenseRepository repository.IExpenseRepository
}

func NewTransactionService(r repository.IRevenueRepository, e repository.IExpenseRepository) *TransactionService {
	return &TransactionService{revenueRepository: r, expenseRepository: e}
}

func (ts TransactionService) AddExpense(dto dtos.ExpenseDto) (error, string) {
	model := dto.ToModel()
	valid, errString := model.Isvalid()
	if !valid {
		return errors.New(errString), "Validation"
	}

	err := ts.expenseRepository.Create(model)
	if err != nil {
		return err, "Internal"
	}

	return nil, ""
}

func (ts TransactionService) AddRevenue(dto dtos.RevenueDto) (error, string) {
	model := dto.ToModel()
	valid, errString := model.Isvalid()
	if !valid {
		return errors.New(errString), "Validation"
	}

	err := ts.revenueRepository.Create(model)
	if err != nil {
		return err, "Internal"
	}

	return nil, ""
}

func (ts TransactionService) GetByIdExpense(id int64) (dtos.ExpenseDto, error, string) {
	expenseModel, err := ts.expenseRepository.GetById(id)
	if err != nil {
		return dtos.ExpenseDto{}, err, "Internal"
	}
	expenseDto := expenseDtoFromModel(expenseModel)
	return expenseDto, nil, ""
}
func (ts TransactionService) GetByIdUserExpense(idUser int64) ([]dtos.ExpenseDto, error, string) {
	expenseModels, err := ts.expenseRepository.GetByUserId(idUser)
	if err != nil {
		return nil, err, "Internal"
	}
	var expensesDtos []dtos.ExpenseDto
	for _, expenseModel := range expenseModels {
		expenseDto := expenseDtoFromModel(expenseModel)
		expensesDtos = append(expensesDtos, expenseDto)
	}
	return expensesDtos, nil, ""
}
func (ts TransactionService) GetByIdRevenue(id int64) (dtos.RevenueDto, error, string) {
	revenueModel, err := ts.revenueRepository.GetById(id)
	if err != nil {
		return dtos.RevenueDto{}, err, "Internal"
	}
	revenueDto := revenueDtoFromModel(revenueModel)
	return revenueDto, nil, ""
}
func (ts TransactionService) GetByIdUserRevenue(idUser int64) ([]dtos.RevenueDto, error, string) {

	revenueModels, err := ts.revenueRepository.GetByUserId(idUser)
	if err != nil {
		return nil, err, "Internal"
	}
	var revenues []dtos.RevenueDto
	for _, revenueModel := range revenueModels {

		revenueDto := revenueDtoFromModel(revenueModel)
		revenues = append(revenues, revenueDto)
	}
	return revenues, nil, ""
}

func (ts TransactionService) UpdateExpense(id int64, dto dtos.ExpenseDto) (error, string) {
	model := dto.ToModel()
	model.ID = id
	valid, errString := model.Isvalid()
	if !valid {
		return errors.New(errString), "Validation"
	}
	err := ts.expenseRepository.Update(model)
	if err != nil {
		return err, "Internal"
	}

	return nil, ""
}
func (ts TransactionService) UpdateRevenue(id int64, dto dtos.RevenueDto) (error, string) {
	model := dto.ToModel()
	model.ID = id
	valid, errString := model.Isvalid()
	if !valid {
		return errors.New(errString), "Validation"
	}
	err := ts.revenueRepository.Update(model)
	if err != nil {
		return err, "Internal"
	}

	return nil, ""
}

func (ts TransactionService) DeleteExpense(id int64, userId int64) (error, string) {

	err := ts.expenseRepository.Delete(id, userId)
	if err != nil {
		return err, "Internal"
	}
	return nil, ""
}
func (ts TransactionService) DeleteRevenue(id int64, userId int64) (error, string) {

	err := ts.revenueRepository.Delete(id, userId)
	if err != nil {
		return err, "Internal"
	}
	return nil, ""
}

func expenseDtoFromModel(ex models.Expense) dtos.ExpenseDto {
	return dtos.ExpenseDto{
		ID:            ex.ID,
		UserId:        ex.UserId,
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
func revenueDtoFromModel(re models.Revenue) dtos.RevenueDto {
	return dtos.RevenueDto{
		ID:          re.ID,
		UserId:      re.UserId,
		Description: re.Description,
		Origin:      re.Origin,
		Type:        re.Type,
		ReceiveDate: re.ReceiveDate,
		IsRecieved:  re.IsRecieved,
		Amount:      re.Amount.ToFloat(),
	}
}
