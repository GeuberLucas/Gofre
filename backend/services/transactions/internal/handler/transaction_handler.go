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
		id, erro := strconv.ParseInt(params["idTransaction"], 10, 64)
		if erro != nil {
			response.ErrorResponse(w, http.StatusBadRequest, erro)
			return
		}
		serviceresult, erro, typeError := s.GetByIdExpense(id)
		if erro != nil {
			if typeError == "Validation" {
				response.ErrorResponse(w, http.StatusBadRequest, erro)
				return
			}
			if typeError == "Internal" {
				response.ErrorResponse(w, http.StatusInternalServerError, erro)
				return
			}

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

		serviceresult, err, typeError := s.GetByIdExpense(int64(userIdToken))
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
		id, erro := strconv.ParseInt(params["idTransaction"], 10, 64)
		if erro != nil {
			response.ErrorResponse(w, http.StatusBadRequest, erro)
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
		erro, typeError := s.UpdateExpense(id, expenseDto)
		if erro != nil {
			if typeError == "Validation" {
				response.ErrorResponse(w, http.StatusBadRequest, erro)
				return
			}
			if typeError == "Internal" {
				response.ErrorResponse(w, http.StatusInternalServerError, erro)
				return
			}

		}

		response.JSONResponse(w, http.StatusOK, nil)
	}
}
func DeleteExpenseHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, erro := strconv.ParseInt(params["idTransaction"], 10, 64)
		if erro != nil {
			response.ErrorResponse(w, http.StatusBadRequest, erro)
			return
		}
		userIdToken, err := jwtToken.ExtractUserId(r)
		if err != nil {
			response.ErrorResponse(w, http.StatusUnauthorized, err)
			return
		}
		erro, typeError := s.DeleteExpense(id, int64(userIdToken))
		if erro != nil {
			if typeError == "Validation" {
				response.ErrorResponse(w, http.StatusBadRequest, erro)
				return
			}
			if typeError == "Internal" {
				response.ErrorResponse(w, http.StatusInternalServerError, erro)
				return
			}

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
		id, erro := strconv.ParseInt(params["idTransaction"], 10, 64)
		if erro != nil {
			response.ErrorResponse(w, http.StatusBadRequest, erro)
			return
		}
		serviceresult, erro, typeError := s.GetByIdRevenue(id)
		if erro != nil {
			if typeError == "Validation" {
				response.ErrorResponse(w, http.StatusBadRequest, erro)
				return
			}
			if typeError == "Internal" {
				response.ErrorResponse(w, http.StatusInternalServerError, erro)
				return
			}

		}

		response.JSONResponse(w, http.StatusOK, serviceresult)
	}
}
func GetByIdUserRevenueHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdToken, err := jwtToken.ExtractUserId(r)
		if err != nil {
			response.ErrorResponse(w, http.StatusUnauthorized, err)
			return
		}
		serviceresult, erro, typeError := s.GetByIdUserRevenue(int64(userIdToken))
		if erro != nil {
			if typeError == "Validation" {
				response.ErrorResponse(w, http.StatusBadRequest, erro)
				return
			}
			if typeError == "Internal" {
				response.ErrorResponse(w, http.StatusInternalServerError, erro)
				return
			}

		}

		response.JSONResponse(w, http.StatusOK, serviceresult)
	}
}
func UpdateRevenueHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, erro := strconv.ParseInt(params["idTransaction"], 10, 64)
		if erro != nil {
			response.ErrorResponse(w, http.StatusBadRequest, erro)
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
		var revenueDto dtos.RevenueDto
		revenueDto.UserId = int64(userIdToken)
		if err = json.Unmarshal(bodyRequest, &revenueDto); err != nil {
			checkErroType(w, err, "validation")
			return
		}
		erro, typeError := s.UpdateRevenue(id, revenueDto)
		if erro != nil {
			if typeError == "Validation" {
				response.ErrorResponse(w, http.StatusBadRequest, erro)
				return
			}
			if typeError == "Internal" {
				response.ErrorResponse(w, http.StatusInternalServerError, erro)
				return
			}

		}

		response.JSONResponse(w, http.StatusOK, nil)
	}
}
func DeleteRevenueHandler(s *service.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, erro := strconv.ParseInt(params["idTransaction"], 10, 64)
		if erro != nil {
			response.ErrorResponse(w, http.StatusBadRequest, erro)
			return
		}
		userIdToken, err := jwtToken.ExtractUserId(r)
		if err != nil {
			response.ErrorResponse(w, http.StatusUnauthorized, err)
			return
		}
		erro, typeError := s.DeleteRevenue(id, int64(userIdToken))
		if erro != nil {
			if typeError == "Validation" {
				response.ErrorResponse(w, http.StatusBadRequest, erro)
				return
			}
			if typeError == "Internal" {
				response.ErrorResponse(w, http.StatusInternalServerError, erro)
				return
			}

		}

		response.JSONResponse(w, http.StatusOK, nil)
	}
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
