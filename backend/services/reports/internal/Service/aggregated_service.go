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

type IAggregatedService interface {
	RegisterEventExpense(tx *sql.Tx, subscriberDTO messaging.MessagingDto) error
	InsertOrUpdateRevenue(model models.Expense) error
	InsertOrUpdateInvestment(model models.Expense) error
}

type AggregatedService struct {
	aggregatedRepository interfaces.IReportsRepository[models.Aggregated]

	revenueRepository    interfaces.IReportsRepository[models.Revenue]
	investmentRepository interfaces.IReportsRepository[models.Investment]
}

// InsertOrUpdateExpense implements IService.

// InsertOrUpdateInvestment implements IService.
func (s *AggregatedService) InsertOrUpdateInvestment(model models.Expense) error {
	panic("unimplemented")
}

// InsertOrUpdateRevenue implements IService.
func (s *AggregatedService) InsertOrUpdateRevenue(model models.Expense) error {
	panic("unimplemented")
}

func NewService(ag interfaces.IReportsRepository[models.Aggregated], rv interfaces.IReportsRepository[models.Revenue], ivt interfaces.IReportsRepository[models.Investment]) IAggregatedService {
	return &AggregatedService{
		aggregatedRepository: ag,
		revenueRepository:    rv,
		investmentRepository: ivt,
	}
}

func (s *AggregatedService) RegisterEventExpense(tx *sql.Tx, subscriberDTO messaging.MessagingDto) error {

	model, _, err := s.aggregatedRepository.GetByMonthAndYear(subscriberDTO.UserId, int(subscriberDTO.Month), int(subscriberDTO.Year))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil

			model = models.Aggregated{
				UserId: subscriberDTO.UserId,
				Month:  int(subscriberDTO.Month),
				Year:   int(subscriberDTO.Year),
			}
		} else {
			return err
		}
	}
	delta := calculateValueForModel(subscriberDTO)
	model.Expense += delta
	switch subscriberDTO.MovementType {
	case string(helpers.ExpenseTypeMensal):
		if subscriberDTO.WithCredit {
			model.MonthlyWithCredit += delta
		} else {
			model.MonthlyWithoutCredit += delta
		}
	case string(helpers.ExpenseTypeVariavel):
		if subscriberDTO.WithCredit {
			model.VariableWithCredit += delta
		} else {
			model.VariableWithoutCredit += delta
		}
	case string(helpers.ExpenseTypeFatura):
		model.Invoice += delta
	}

	model.Result = model.Revenue - model.Expense - model.Investments
	_, err = s.aggregatedRepository.InsertOrUpdate(tx, &model)
	if err != nil {
		return err
	}
	return nil
}
func calculateValueForModel(subscriberDTO messaging.MessagingDto) types.Money {
	switch subscriberDTO.Action {
	case messaging.ActionUpdate:
		return subscriberDTO.Amount - subscriberDTO.AmountOld
	case messaging.ActionDelete:
		return -subscriberDTO.Amount
	case messaging.ActionInsert:
		return subscriberDTO.Amount
	}
	return types.Money(0)
}
