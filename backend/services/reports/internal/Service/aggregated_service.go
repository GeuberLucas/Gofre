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
	RegisterEventRevenue(tx *sql.Tx, subscriberDTO messaging.MessagingDto) error
	RegisterEventInvestment(tx *sql.Tx, subscriberDTO messaging.MessagingDto) error
}

type AggregatedService struct {
	aggregatedRepository interfaces.IAggregatedRepository

	// revenueRepository    interfaces.IReportsRepository[models.Revenue]
	// investmentRepository interfaces.IReportsRepository[models.Investment]
}

func (s *AggregatedService) RegisterEventRevenue(
	tx *sql.Tx,
	subscriberDTO messaging.MessagingDto,
) error {
	return s.RegisterEvent(tx, subscriberDTO, s.HasBasicChanges)
}

func (s *AggregatedService) RegisterEventInvestment(
	tx *sql.Tx,
	subscriberDTO messaging.MessagingDto,
) error {

	return s.RegisterEvent(tx, subscriberDTO, s.HasBasicChanges)
}

func (s *AggregatedService) RegisterEvent(
	tx *sql.Tx,
	subscriberDTO messaging.MessagingDto,
	hasDiff func(messaging.MessagingDto) bool,
) error {
	if err := subscriberDTO.IsValid(); err != nil {
		return err
	}
	model := models.Aggregated{
		UserId: subscriberDTO.UserId,
		Month:  int(subscriberDTO.Month),
		Year:   int(subscriberDTO.Year),
	}
	switch subscriberDTO.Action {
	case messaging.ActionInsert:
		if subscriberDTO.IsConfirmed {
			return s.process(tx, &subscriberDTO, &model, 1)
		}

	case messaging.ActionDelete:
		if subscriberDTO.IsConfirmed {
			return s.process(tx, &subscriberDTO, &model, -1)
		}

	case messaging.ActionUpdate:
		if !hasDiff(subscriberDTO) {
			return nil
		}
		oldModel, oldSub := s.createOldStructs(subscriberDTO)
		if subscriberDTO.IsConfirmedOld {
			if err := s.process(tx, &oldSub, &oldModel, -1); err != nil {
				return err
			}
		}
		if subscriberDTO.IsConfirmed {
			if err := s.process(tx, &subscriberDTO, &model, 1); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewService(ag interfaces.IReportsRepository[models.Aggregated],

// rv interfaces.IReportsRepository[models.Revenue],
// ivt interfaces.IReportsRepository[models.Investment],
) *AggregatedService {
	return &AggregatedService{
		aggregatedRepository: ag,
		// revenueRepository:    rv,
		// investmentRepository: ivt,
	}
}

func (s *AggregatedService) RegisterEventExpense(
	tx *sql.Tx,
	subscriberDTO messaging.MessagingDto,
) error {
	return s.RegisterEvent(tx, subscriberDTO, s.HasDetailedChanges)
}

func (s *AggregatedService) HasDetailedChanges(subscriberDTO messaging.MessagingDto) bool {
	diffAmount := subscriberDTO.AmountOld != subscriberDTO.Amount
	diffType := subscriberDTO.MovementTypeOld != subscriberDTO.MovementType
	diffMonth := subscriberDTO.Month != subscriberDTO.MonthOld
	diffYear := subscriberDTO.YearOld != subscriberDTO.Year
	diffCredit := subscriberDTO.WithCredit != subscriberDTO.WithCreditOld
	diffConfirmed := subscriberDTO.IsConfirmed != subscriberDTO.IsConfirmedOld
	return diffAmount || diffType || diffMonth || diffYear || diffCredit || diffConfirmed
}
func (s *AggregatedService) HasBasicChanges(subscriberDTO messaging.MessagingDto) bool {
	diffAmount := subscriberDTO.AmountOld != subscriberDTO.Amount
	diffMonth := subscriberDTO.Month != subscriberDTO.MonthOld
	diffYear := subscriberDTO.YearOld != subscriberDTO.Year
	diffConfirmed := subscriberDTO.IsConfirmed != subscriberDTO.IsConfirmedOld
	return diffAmount || diffMonth || diffYear || diffConfirmed
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
func (s *AggregatedService) process(tx *sql.Tx, dto *messaging.MessagingDto, model *models.Aggregated, multiplier int) error {
	existingRecord, _, err := s.aggregatedRepository.GetByMonthAndYear(dto.UserId, int(dto.Month), int(dto.Year))
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	if err == nil {
		*model = existingRecord
	}
	calculateValuesModel(dto, model, multiplier)
	model.Result = model.Revenue - model.Expense - model.Investments
	_, err = s.aggregatedRepository.InsertOrUpdate(tx, model)
	if err != nil {
		return err
	}
	return nil
}
func calculateValuesModel(dto *messaging.MessagingDto, model *models.Aggregated, multiplier int) {
	switch dto.Movement {
	case messaging.TypeExpense:
		calculateValueForExpense(dto, model, multiplier)
	case messaging.TypeIncome:
		delta := dto.Amount * types.Money(multiplier)
		model.Revenue += delta
	case messaging.TypeInvestment:
		delta := dto.Amount * types.Money(multiplier)
		model.Investments += delta
	}
}
func calculateValueForExpense(subscriberDTO *messaging.MessagingDto, model *models.Aggregated, multiplier int) {
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
