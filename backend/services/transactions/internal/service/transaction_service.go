package service

import (
	"encoding/json"
	"fmt"

	"github.com/GeuberLucas/Gofre/backend/pkg/helpers"
	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
	dtos "github.com/GeuberLucas/Gofre/backend/services/transaction/internal/Dtos"
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/models"
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/repository"
)

type TransactionService struct {
	revenueRepository repository.IRevenueRepository
	expenseRepository repository.IExpenseRepository
	broker            messaging.IMessaging
}

func NewTransactionService(r repository.IRevenueRepository, e repository.IExpenseRepository, b messaging.IMessaging) *TransactionService {
	return &TransactionService{revenueRepository: r, expenseRepository: e, broker: b}
}

func (ts *TransactionService) AddExpense(dto dtos.ExpenseDto) (string, error) {
	model := dto.ToModel()
	err := model.Isvalid()
	if err != nil {
		return "Validation", err
	}

	err = ts.expenseRepository.Create(model)
	if err != nil {
		return "Internal", err
	}
	err = ts.sendExpenseToBroker(&model, nil, messaging.ActionInsert)
	if err != nil {
		return helpers.INTERNAL.String(), err
	}
	return "", nil
}

func (ts *TransactionService) AddRevenue(dto dtos.RevenueDto) (string, error) {
	model := dto.ToModel()
	err := model.Isvalid()
	if err != nil {
		return "Validation", err
	}

	err = ts.revenueRepository.Create(model)
	if err != nil {
		return "Internal", err
	}
	err = ts.sendRevenueToBroker(&model, nil, messaging.ActionInsert)
	if err != nil {
		return helpers.INTERNAL.String(), err
	}
	return "", err
}

func (ts *TransactionService) GetByIdExpense(id int64) (dtos.ExpenseDto, error, string) {
	expenseModel, err := ts.expenseRepository.GetById(id)
	if err != nil {
		return dtos.ExpenseDto{}, err, "Internal"
	}
	expenseDto := expenseDtoFromModel(expenseModel)
	return expenseDto, nil, ""
}
func (ts *TransactionService) GetByIdUserExpense(idUser int64) ([]dtos.ExpenseDto, error, string) {
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
func (ts *TransactionService) GetByIdRevenue(id int64) (dtos.RevenueDto, error, string) {
	revenueModel, err := ts.revenueRepository.GetById(id)
	if err != nil {
		return dtos.RevenueDto{}, err, "Internal"
	}
	revenueDto := revenueDtoFromModel(revenueModel)
	return revenueDto, nil, ""
}
func (ts *TransactionService) GetByIdUserRevenue(idUser int64) ([]dtos.RevenueDto, error, string) {

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

func (ts *TransactionService) UpdateExpense(id int64, dto dtos.ExpenseDto) (error, string) {
	oldModel, err := ts.expenseRepository.GetById(id)
	if err != nil {
		return err, "Internal"
	}
	model := dto.ToModel()
	model.ID = id
	err = model.Isvalid()
	if err != nil {
		return err, "Validation"
	}
	err = ts.expenseRepository.Update(model)
	if err != nil {
		return err, "Internal"
	}
	err = ts.sendExpenseToBroker(&model, &oldModel, messaging.ActionUpdate)
	if err != nil {
		return err, helpers.INTERNAL.String()
	}
	return nil, ""
}
func (ts *TransactionService) UpdateRevenue(id int64, dto dtos.RevenueDto) (error, string) {
	oldModel, err := ts.revenueRepository.GetById(id)
	if err != nil {
		return err, "Internal"
	}
	model := dto.ToModel()
	model.ID = id
	err = model.Isvalid()
	if err != nil {
		return err, "Validation"
	}
	err = ts.revenueRepository.Update(model)
	if err != nil {
		return err, "Internal"
	}
	err = ts.sendRevenueToBroker(&model, &oldModel, messaging.ActionUpdate)
	if err != nil {
		return err, helpers.INTERNAL.String()
	}
	return nil, ""
}

func (ts *TransactionService) DeleteExpense(id int64, userId int64) (error, string) {
	model, err := ts.expenseRepository.GetById(id)
	if err != nil {
		return err, "Internal"
	}
	err = ts.expenseRepository.Delete(id, userId)
	if err != nil {
		return err, "Internal"
	}
	err = ts.sendExpenseToBroker(&model, nil, messaging.ActionDelete)
	if err != nil {
		return err, helpers.INTERNAL.String()
	}
	return nil, ""
}
func (ts *TransactionService) DeleteRevenue(id int64, userId int64) (error, string) {
	model, err := ts.revenueRepository.GetById(id)
	if err != nil {
		return err, "Internal"
	}
	err = ts.revenueRepository.Delete(id, userId)
	if err != nil {
		return err, "Internal"
	}

	err = ts.sendRevenueToBroker(&model, nil, messaging.ActionDelete)
	if err != nil {
		return err, helpers.INTERNAL.String()
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
func (ts *TransactionService) sendRevenueToBroker(model *models.Revenue, modelOld *models.Revenue, action messaging.ActionType) error {

	ms := messaging.MessagingDto{
		Month:            model.ReceiveDate.Month(),
		Year:             uint(model.ReceiveDate.Local().Year()),
		Amount:           model.Amount,
		Movement:         messaging.TypeIncome,
		MovementType:     model.Type,
		MovementCategory: "",
		WithCredit:       false,
		IsConfirmed:      model.IsRecieved,
		Action:           action,
		UserId:           int(model.UserId),
	}

	if action == messaging.ActionUpdate {
		ms.AmountOld = modelOld.Amount
		ms.MonthOld = modelOld.ReceiveDate.Month()
		ms.YearOld = uint(modelOld.ReceiveDate.Year())
		ms.MovementCategoryOld = ""
		ms.MovementTypeOld = modelOld.Type
		ms.WithCreditOld = false
		ms.IsConfirmedOld = modelOld.IsRecieved
	}

	if err := ms.IsValid(); err != nil {
		return err
	}

	eventName := fmt.Sprintf("finance.%s.%s", messaging.TypeInvestment, action)
	json, err := json.Marshal(ms)
	if err != nil {
		return err
	}
	ts.broker.PublishMessage(eventName, json)

	return nil

}
func (ts *TransactionService) sendExpenseToBroker(model *models.Expense, modelOld *models.Expense, action messaging.ActionType) error {

	ms := messaging.MessagingDto{
		Month:            model.PaymentDate.Month(),
		Year:             uint(model.PaymentDate.Local().Year()),
		Amount:           model.Amount,
		Movement:         messaging.TypeExpense,
		MovementType:     string(model.Type),
		MovementCategory: string(model.Category),
		WithCredit:       model.PaymentMethod == helpers.PaymentMethodCredito,
		IsConfirmed:      model.IsPaid,
		Action:           action,
		UserId:           int(model.UserId),
	}

	if action == messaging.ActionUpdate {
		ms.AmountOld = modelOld.Amount
		ms.MonthOld = modelOld.PaymentDate.Month()
		ms.YearOld = uint(modelOld.PaymentDate.Year())
		ms.MovementCategoryOld = string(modelOld.Category)
		ms.MovementTypeOld = string(modelOld.Type)
		ms.WithCreditOld = modelOld.PaymentMethod == helpers.PaymentMethodCredito
		ms.IsConfirmedOld = modelOld.IsPaid
	}

	if err := ms.IsValid(); err != nil {
		return err
	}

	eventName := fmt.Sprintf("finance.%s.%s", messaging.TypeInvestment, action)
	json, err := json.Marshal(ms)
	if err != nil {
		return err
	}
	ts.broker.PublishMessage(eventName, json)

	return nil

}
