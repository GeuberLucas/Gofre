package service

import (
	"database/sql"
	"errors"

	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
	"github.com/GeuberLucas/Gofre/backend/pkg/types"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/interfaces"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/models"
)

type IExpenseService interface {
	RegisterEvent(subscriberDTO messaging.MessagingDto) error
}

type ExpenseService struct {
	expenseRepository interfaces.IReportsRepository[models.Expense]
}

func NewExpenseService(repo interfaces.IReportsRepository[models.Expense]) IExpenseService {
	return &ExpenseService{
		expenseRepository: repo,
	}
}

func (s *ExpenseService) InsertExpense(subscriberDTO messaging.MessagingDto) error {

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
	_, err = s.expenseRepository.InsertOrUpdate(&model)
	if err != nil {
		return err
	}
	return nil
}
func (s *ExpenseService) UpdateExpense(subscriberDTO messaging.MessagingDto) error {
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
	_, err = s.expenseRepository.InsertOrUpdate(&model)
	if err != nil {
		return err
	}
	return nil
}
func (s *ExpenseService) DeleteExpense(subscriberDTO messaging.MessagingDto) error {
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
	_, err = s.expenseRepository.InsertOrUpdate(&model)
	if err != nil {
		return err
	}
	return nil
}
func calculateValuesExpense(dto messaging.MessagingDto, model *models.Expense, amount types.Money) error {
	//tipo de gasto
	switch dto.MovementType {
	case "Fatura":
		model.Invoice += amount
	case "Mensal":
		model.Monthly += amount
	case "Variavel":
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

func (s *ExpenseService) RegisterEvent(subscriberDTO messaging.MessagingDto) error {
	switch subscriberDTO.Action {
	case messaging.ActionInsert:
		err := s.InsertExpense(subscriberDTO)
		return err
	case messaging.ActionUpdate:
		err := s.UpdateExpense(subscriberDTO)
		return err
	case messaging.ActionDelete:
		err := s.DeleteExpense(subscriberDTO)
		return err
	}
	return nil
}
