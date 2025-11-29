package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	jwtToken "github.com/GeuberLucas/Gofre/backend/pkg/jwt"
	"github.com/GeuberLucas/Gofre/backend/pkg/response"
	dtos "github.com/GeuberLucas/Gofre/backend/services/transaction/internal/Dtos"
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/service"
	"github.com/gorilla/mux"
)

func AddExpenseHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyRequest, err := io.ReadAll(r.Body)
		if err != nil {
			response.ErrorResponse(w, http.StatusUnprocessableEntity, err)
			return
		}
		userIdToken, err := jwtToken.ExtractUserId(r)
		if err != nil {
			response.ErrorResponse(w, http.StatusUnauthorized, err)
			return
		}
		var expenseDto dtos.ExpenseDto
		if err = json.Unmarshal(bodyRequest, &expenseDto); err != nil {
			checkErroType(w, err, "validation")
			return
		}
		expenseDto.UserId = int64(userIdToken)
		err, stringTypeError := s.AddExpense(expenseDto)
		if err != nil {
			checkErroType(w, err, stringTypeError)
			return
		}

	}
}
func GetByIdExpenseHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseInt(params["idTransaction"], 10, 64)
		if err != nil {
			response.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}
		serviceresult, err, typeError := s.GetByIdExpense(id)
		if err != nil {
			checkErroType(w, err, typeError)
			return

		}

		response.JSONResponse(w, http.StatusOK, serviceresult)
	}
}
func GetByIdUserExpenseHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userIdToken, err := jwtToken.ExtractUserId(r)
		if err != nil {
			response.ErrorResponse(w, http.StatusUnauthorized, err)
			return
		}

		serviceresult, err, typeError := s.GetByIdUserExpense(int64(userIdToken))
		if err != nil {
			if typeError == "Validation" {
				response.ErrorResponse(w, http.StatusBadRequest, err)
				return
			}
			if typeError == "Internal" {
				response.ErrorResponse(w, http.StatusInternalServerError, err)
				return
			}

		}

		response.JSONResponse(w, http.StatusOK, serviceresult)
	}
}
func UpdateExpenseHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseInt(params["idTransaction"], 10, 64)
		if err != nil {
			response.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}
		userIdToken, err := jwtToken.ExtractUserId(r)
		if err != nil {
			response.ErrorResponse(w, http.StatusUnauthorized, err)
			return
		}
		bodyRequest, err := io.ReadAll(r.Body)
		if err != nil {
			response.ErrorResponse(w, http.StatusUnprocessableEntity, err)
			return
		}
		var expenseDto dtos.ExpenseDto
		expenseDto.UserId = int64(userIdToken)
		if err = json.Unmarshal(bodyRequest, &expenseDto); err != nil {
			checkErroType(w, err, "validation")
			return
		}
		err, typeError := s.UpdateExpense(id, expenseDto)
		if err != nil {
			checkErroType(w, err, typeError)
			return

		}

		response.JSONResponse(w, http.StatusOK, nil)
	}
}
func DeleteExpenseHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseInt(params["idTransaction"], 10, 64)
		if err != nil {
			checkErroType(w, err, "Validation")
			return
		}
		userIdToken, err := jwtToken.ExtractUserId(r)
		if err != nil {
			checkErroType(w, err, "Validation")
			return
		}
		err, typeError := s.DeleteExpense(id, int64(userIdToken))
		if err != nil {
			checkErroType(w, err, typeError)
			return

		}

		response.JSONResponse(w, http.StatusOK, nil)
	}
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
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseInt(params["idTransaction"], 10, 64)
		if err != nil {
			checkErroType(w, err, "Validation")
			return
		}
		serviceresult, err, typeError := s.GetByIdRevenue(id)
		if err != nil {
			checkErroType(w, err, typeError)
			return

		}

		response.JSONResponse(w, http.StatusOK, serviceresult)
	}
}
func GetByIdUserRevenueHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdToken, err := jwtToken.ExtractUserId(r)
		if err != nil {
			checkErroType(w, err, "Validation")
			return
		}
		serviceresult, err, typeError := s.GetByIdUserRevenue(int64(userIdToken))
		if err != nil {
			checkErroType(w, err, typeError)
			return

		}

		response.JSONResponse(w, http.StatusOK, serviceresult)
	}
}
func UpdateRevenueHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseInt(params["idTransaction"], 10, 64)
		if err != nil {
			checkErroType(w, err, "Validation")
			return
		}
		userIdToken, err := jwtToken.ExtractUserId(r)
		if err != nil {
			checkErroType(w, err, "Validation")
			return
		}
		bodyRequest, err := io.ReadAll(r.Body)
		if err != nil {
			response.ErrorResponse(w, http.StatusUnprocessableEntity, err)
			return
		}
		var revenueDto dtos.RevenueDto
		revenueDto.UserId = int64(userIdToken)
		if err = json.Unmarshal(bodyRequest, &revenueDto); err != nil {
			checkErroType(w, err, "Validation")
			return
		}
		err, typeError := s.UpdateRevenue(id, revenueDto)
		if err != nil {
			checkErroType(w, err, typeError)
			return

		}

		response.JSONResponse(w, http.StatusOK, nil)
	}
}
func DeleteRevenueHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseInt(params["idTransaction"], 10, 64)
		if err != nil {
			response.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}
		userIdToken, err := jwtToken.ExtractUserId(r)
		if err != nil {
			checkErroType(w, err, "Validation")
			return
		}
		err, typeError := s.DeleteRevenue(id, int64(userIdToken))
		if err != nil {

			checkErroType(w, err, typeError)
			return
		}

		response.JSONResponse(w, http.StatusOK, nil)
	}
}

func checkErroType(w http.ResponseWriter, err error, typeError string) {
	switch typeError {
	case "Validation":
		response.ErrorResponse(w, http.StatusBadRequest, err)
		return
	default:
		response.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

}
