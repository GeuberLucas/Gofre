package service

import (
	"errors"

	dtos "github.com/GeuberLucas/Gofre/backend/services/transaction/internal/Dtos"
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
