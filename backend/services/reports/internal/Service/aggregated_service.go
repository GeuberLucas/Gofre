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

	model := models.Aggregated{
		UserId: subscriberDTO.UserId,
		Month:  int(subscriberDTO.Month),
		Year:   int(subscriberDTO.Year),
	}
	switch subscriberDTO.Action {
	case messaging.ActionInsert:
		if subscriberDTO.IsConfirmed {
			err := s.processExpense(tx, &subscriberDTO, &model, 1)
			return err
		}
	case messaging.ActionDelete:
		if subscriberDTO.IsConfirmed {
			return s.processExpense(tx, &subscriberDTO, &model, -1)
		}
	case messaging.ActionUpdate:
		existsDiff := s.DetermineDiff(subscriberDTO)
		if existsDiff {
			oldModel, oldSub := s.createOldStructs(subscriberDTO)
			if subscriberDTO.IsConfirmedOld {
				err := s.processExpense(tx, &oldSub, &oldModel, -1)
				if err != nil {
					return err
				}
			}
			if subscriberDTO.IsConfirmed {
				err := s.processExpense(tx, &subscriberDTO, &model, 1)
				if err != nil {
					return err
				}
			}
			return nil
		}
	}

	return nil
}
func (s *AggregatedService) DetermineDiff(subscriberDTO messaging.MessagingDto) bool {
	diffAmount := subscriberDTO.AmountOld != subscriberDTO.Amount
	diffType := subscriberDTO.MovementTypeOld != subscriberDTO.MovementType
	diffMonth := subscriberDTO.Month != subscriberDTO.MonthOld
	diffYear := subscriberDTO.YearOld != subscriberDTO.Year
	diffCredit := subscriberDTO.WithCredit != subscriberDTO.WithCreditOld
	diffConfirmed := subscriberDTO.IsConfirmed != subscriberDTO.IsConfirmedOld
	return diffAmount || diffType || diffMonth || diffYear || diffCredit || diffConfirmed
}
func (s *AggregatedService) createOldStructs(subscriberDTO messaging.MessagingDto) (models.Aggregated, messaging.MessagingDto) {
	oldModel := models.Aggregated{
		UserId: subscriberDTO.UserId,
		Month:  int(subscriberDTO.MonthOld),
		Year:   int(subscriberDTO.YearOld),
	}
	oldSub := messaging.MessagingDto{
		Amount:         subscriberDTO.AmountOld,
		MovementType:   subscriberDTO.MovementTypeOld,
		IsConfirmed:    subscriberDTO.IsConfirmedOld,
		Month:          subscriberDTO.MonthOld,
		Year:           subscriberDTO.YearOld,
		WithCredit:     subscriberDTO.WithCreditOld,
		IsConfirmedOld: subscriberDTO.IsConfirmedOld,
	}
	return oldModel, oldSub
}
func (s *AggregatedService) processExpense(tx *sql.Tx, dto *messaging.MessagingDto, model *models.Aggregated, multiplier int) error {
	existingRecord, _, err := s.aggregatedRepository.GetByMonthAndYear(dto.UserId, int(dto.Month), int(dto.Year))
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	if err == nil {

		*model = existingRecord
	}
	calculateValueForModel(dto, model, multiplier)
	model.Result = model.Revenue - model.Expense - model.Investments
	_, err = s.aggregatedRepository.InsertOrUpdate(tx, model)
	if err != nil {
		return err
	}
	return nil
}

func calculateValueForModel(subscriberDTO *messaging.MessagingDto, model *models.Aggregated, multiplier int) {
	delta := subscriberDTO.Amount * types.Money(multiplier)
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
}
