package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/GeuberLucas/Gofre/backend/pkg/response"
	dtos "github.com/GeuberLucas/Gofre/backend/services/transaction/internal/Dtos"
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/service"
)

func AddExpenseHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyRequest, err := io.ReadAll(r.Body)
		if err != nil {
			response.ErrorResponse(w, http.StatusUnprocessableEntity, err)
			return
		}
		var expenseDto dtos.ExpenseDto
		if err = json.Unmarshal(bodyRequest, &expenseDto); err != nil {
			checkErroType(w, err, "validation")
			return
		}

		err, stringTypeError := s.AddExpense(expenseDto)
		if err != nil {
			checkErroType(w, err, stringTypeError)
			return
		}

	}
}
func GetByIdExpenseHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func GetByIdUserExpenseHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func UpdateExpenseHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func DeleteExpenseHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func AddRevenueHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyRequest, err := io.ReadAll(r.Body)
		if err != nil {
			response.ErrorResponse(w, http.StatusUnprocessableEntity, err)
			return
		}
		var revenueDto dtos.RevenueDto
		if err = json.Unmarshal(bodyRequest, &revenueDto); err != nil {
			checkErroType(w, err, "validation")
			return
		}

		err, stringTypeError := s.AddRevenue(revenueDto)
		if err != nil {
			checkErroType(w, err, stringTypeError)
			return
		}

	}
}
func GetByIdRevenueHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func GetByIdUserRevenueHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func UpdateRevenueHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func DeleteRevenueHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func checkErroType(w http.ResponseWriter, err error, typeError string) {
	switch typeError {
	case "validation":
		response.ErrorResponse(w, http.StatusBadRequest, err)
		return
	default:
		response.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

}
