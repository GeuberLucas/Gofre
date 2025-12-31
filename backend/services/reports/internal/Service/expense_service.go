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

	err = calculateValuesExpense(subscriberDTO, &model, subscriberDTO.Amount, 1)
	if err != nil {
		return err
	}
	_, err = s.expenseRepository.InsertOrUpdate(tx, &model)
	if err != nil {
		return err
	}
	return nil
}
func (s *ExpenseService) UpdateExpense(tx *sql.Tx, dto messaging.MessagingDto) error {

	oldStateDto := messaging.MessagingDto{
		MovementType:     dto.MovementTypeOld,
		MovementCategory: dto.MovementCategoryOld,
		IsConfirmed:      dto.IsConfirmedOld,
		WithCredit:       dto.WithCreditOld,
	}

	err := s.processExpenseChange(tx, dto.UserId, int(dto.MonthOld), int(dto.YearOld), oldStateDto, dto.AmountOld, -1)
	if err != nil {
		return err
	}

	err = s.processExpenseChange(tx, dto.UserId, int(dto.Month), int(dto.Year), dto, dto.Amount, 1)
	if err != nil {
		return err
	}

	return nil
}

func (s *ExpenseService) processExpenseChange(tx *sql.Tx, userId int, month int, year int, dto messaging.MessagingDto, amount types.Money, multiplier int) error {

	model, _, err := s.expenseRepository.GetByMonthAndYear(userId, month, year)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			model = models.Expense{
				UserId: userId,
				Month:  month,
				Year:   year,
			}
		} else {
			return err
		}
	}

	if err := calculateValuesExpense(dto, &model, amount, multiplier); err != nil {
		return err
	}

	if _, err := s.expenseRepository.InsertOrUpdate(tx, &model); err != nil {
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

	err = calculateValuesExpense(subscriberDTO, &model, subscriberDTO.Amount, -1)
	if err != nil {
		return err
	}
	_, err = s.expenseRepository.InsertOrUpdate(tx, &model)
	if err != nil {
		return err
	}
	return nil
}

func calculateValuesExpense(dto messaging.MessagingDto, model *models.Expense, amount types.Money, multiplier int) error {

	finalAmount := amount * types.Money(multiplier)

	switch dto.MovementType {
	case string(helpers.ExpenseTypeFatura):
		model.Invoice += finalAmount
	case string(helpers.ExpenseTypeMensal):
		model.Monthly += finalAmount
	case string(helpers.ExpenseTypeVariavel):
		model.Variable += finalAmount
	default:
		return errors.New("Type of movement not recognized")
	}

	if !dto.WithCredit {
		model.Planned += finalAmount

		if dto.IsConfirmed {
			model.Actual += finalAmount
		} else {
			model.Pending += finalAmount
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
