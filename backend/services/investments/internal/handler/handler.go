package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GeuberLucas/Gofre/backend/pkg/response"
	dtos "github.com/GeuberLucas/Gofre/backend/services/investments/internal/DTOs"
	"github.com/GeuberLucas/Gofre/backend/services/investments/internal/service"
	"github.com/gorilla/mux"
)

type IHandlerService interface {
	AddInvestmentHandler() http.HandlerFunc
	GetInvestmentHandler() http.HandlerFunc
	GetByIdInvestmentHandler() http.HandlerFunc
	UpdateInvestmentHandler() http.HandlerFunc
	DeleteInvestmentHandler() http.HandlerFunc
}

type HandlerService struct {
	portfolioService service.IPortfolioService
}

func NewHandlerService(service service.IPortfolioService) IHandlerService {
	return &HandlerService{
		portfolioService: service,
	}
}

func (hd *HandlerService) AddInvestmentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto dtos.Portfolio
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&dto); err != nil {
			response.ErrorResponse(w, 500, err)
		}
		userIdHeader := r.Header.Get("user_id")
		userId, err := strconv.ParseInt(userIdHeader, 10, 64)
		if err != nil {
			response.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}
		dto.User_id = int(userId)
		err, stringTypeError := hd.portfolioService.Add(dto)
		if err != nil {
			checkErroType(w, err, stringTypeError)
			return
		}
	}
}
func (hd *HandlerService) GetInvestmentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdHeader := r.Header.Get("user_id")
		userId, err := strconv.ParseInt(userIdHeader, 10, 64)
		if err != nil {
			response.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}
		serviceresult, err, typeError := hd.portfolioService.GetAll(userIdInt)
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
func (hd *HandlerService) GetByIdInvestmentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseInt(params["idTransaction"], 10, 64)
		if err != nil {
			checkErroType(w, err, "Validation")
			return
		}
		serviceresult, err, typeError := hd.portfolioService.GetById(id)
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
func (hd *HandlerService) UpdateInvestmentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseInt(params["idTransaction"], 10, 64)
		if err != nil {
			checkErroType(w, err, "Validation")
			return
		}
		var dto dtos.Portfolio
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&dto); err != nil {
			response.ErrorResponse(w, 500, err)
		}
		userIdHeader := r.Header.Get("user_id")
		userId, err := strconv.ParseInt(userIdHeader, 10, 64)
		if err != nil {
			response.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}
		dto.Id = uint(id)
		dto.User_id = int(userId)
		err, stringTypeError := hd.portfolioService.Update(dto)
		if err != nil {
			checkErroType(w, err, stringTypeError)
			return
		}

	}
}
func (hd *HandlerService) DeleteInvestmentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseInt(params["idTransaction"], 10, 64)
		if err != nil {
			checkErroType(w, err, "Validation")
			return
		}
		err, stringTypeError := hd.portfolioService.Delete(uint(id))
		if err != nil {
			checkErroType(w, err, stringTypeError)
			return
		}
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
