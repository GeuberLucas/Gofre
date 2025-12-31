package service

import (
	"database/sql"
	"errors"

	"github.com/GeuberLucas/Gofre/backend/pkg/helpers"
	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
	"github.com/GeuberLucas/Gofre/backend/pkg/types"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/interfaces"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/models"
)

type IExpenseService interface {
	RegisterEvent(tx *sql.Tx, subscriberDTO messaging.MessagingDto) error
}

type ExpenseService struct {
	expenseRepository interfaces.IReportsRepository[models.Expense]
}

func NewExpenseService(repo interfaces.IReportsRepository[models.Expense]) IExpenseService {
	return &ExpenseService{
		expenseRepository: repo,
	}
}

func (s *ExpenseService) InsertExpense(tx *sql.Tx, subscriberDTO messaging.MessagingDto) error {

	model, _, err := s.expenseRepository.GetByMonthAndYear(subscriberDTO.UserId, int(subscriberDTO.Month), int(subscriberDTO.Year))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil

			model = models.Expense{
				UserId: subscriberDTO.UserId,
				Month:  int(subscriberDTO.Month),
				Year:   int(subscriberDTO.Year),
			}
		} else {
			return err
		}
	}
	var amount = subscriberDTO.Amount

	err = calculateValuesExpense(subscriberDTO, &model, amount)
	if err != nil {
		return err
	}
	_, err = s.expenseRepository.InsertOrUpdate(tx, &model)
	if err != nil {
		return err
	}
	return nil
}
func (s *ExpenseService) UpdateExpense(tx *sql.Tx, subscriberDTO messaging.MessagingDto) error {
	var model models.Expense
	model.Month = int(subscriberDTO.Month)
	model.Year = int(subscriberDTO.Year)
	model.UserId = subscriberDTO.UserId
	model, _, err := s.expenseRepository.GetByMonthAndYear(subscriberDTO.UserId, model.Month, model.Year)
	if err != nil {
		return err
	}
	var amount = subscriberDTO.Amount - subscriberDTO.AmountOld

	err = calculateValuesExpense(subscriberDTO, &model, amount)
	if err != nil {
		return err
	}
	_, err = s.expenseRepository.InsertOrUpdate(tx, &model)
	if err != nil {
		return err
	}
	return nil
}
func (s *ExpenseService) DeleteExpense(tx *sql.Tx, subscriberDTO messaging.MessagingDto) error {
	var model models.Expense
	model.Month = int(subscriberDTO.Month)
	model.Year = int(subscriberDTO.Year)
	model.UserId = subscriberDTO.UserId
	model, _, err := s.expenseRepository.GetByMonthAndYear(subscriberDTO.UserId, model.Month, model.Year)
	if err != nil {
		return err
	}
	var amount = -subscriberDTO.Amount

	err = calculateValuesExpense(subscriberDTO, &model, amount)
	if err != nil {
		return err
	}
	_, err = s.expenseRepository.InsertOrUpdate(tx, &model)
	if err != nil {
		return err
	}
	return nil
}
func calculateValuesExpense(dto messaging.MessagingDto, model *models.Expense, amount types.Money) error {
	//tipo de gasto
	switch dto.MovementType {
	case string(helpers.ExpenseTypeFatura):
		model.Invoice += amount
	case string(helpers.ExpenseTypeMensal):
		model.Monthly += amount
	case string(helpers.ExpenseTypeVariavel):
		model.Variable += amount
	default:
		return errors.New("Type of movement not reconized")
	}

	//meio de pagamento e pago ou nao
	if !dto.WithCredit {
		model.Planned += amount
		if dto.IsConfirmed {
			model.Actual += amount
		} else {
			model.Pending += amount
		}
	}
	return nil
}

func (s *ExpenseService) RegisterEvent(tx *sql.Tx, subscriberDTO messaging.MessagingDto) error {
	switch subscriberDTO.Action {
	case messaging.ActionInsert:
		err := s.InsertExpense(tx, subscriberDTO)
		return err
	case messaging.ActionUpdate:
		err := s.UpdateExpense(tx, subscriberDTO)
		return err
	case messaging.ActionDelete:
		err := s.DeleteExpense(tx, subscriberDTO)
		return err
	}
	return nil
}
