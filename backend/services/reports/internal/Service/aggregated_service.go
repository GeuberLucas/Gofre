package service

import (
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/interfaces"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/models"
)

type IService interface {
	InsertOrUpdateExpense(model models.Expense) error
	InsertOrUpdateRevenue(model models.Expense) error
	InsertOrUpdateInvestment(model models.Expense) error
}

type Service struct {
	aggregatedRepository interfaces.IReportsRepository[models.Aggregated]
	expenseRepository    interfaces.IReportsRepository[models.Expense]
	revenueRepository    interfaces.IReportsRepository[models.Revenue]
	investmentRepository interfaces.IReportsRepository[models.Investment]
}

// InsertOrUpdateExpense implements IService.
func (s *Service) InsertOrUpdateExpense(model models.Expense) error {
	panic("unimplemented")
}

// InsertOrUpdateInvestment implements IService.
func (s *Service) InsertOrUpdateInvestment(model models.Expense) error {
	panic("unimplemented")
}

// InsertOrUpdateRevenue implements IService.
func (s *Service) InsertOrUpdateRevenue(model models.Expense) error {
	panic("unimplemented")
}

func NewService(ag interfaces.IReportsRepository[models.Aggregated], ex interfaces.IReportsRepository[models.Expense], rv interfaces.IReportsRepository[models.Revenue], ivt interfaces.IReportsRepository[models.Investment]) IService {
	return &Service{
		aggregatedRepository: ag,
		revenueRepository:    rv,
		expenseRepository:    ex,
		investmentRepository: ivt,
	}
}
