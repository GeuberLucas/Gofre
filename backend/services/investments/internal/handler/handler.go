package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GeuberLucas/Gofre/backend/pkg/response"
	dtos "github.com/GeuberLucas/Gofre/backend/services/investments/internal/DTOs"
	"github.com/GeuberLucas/Gofre/backend/services/investments/internal/helpers"
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
			checkErroType(w, helpers.INTERNAL, err)
			return
		}
		userIdHeader := r.Header.Get("user_id")
		userId, err := strconv.ParseInt(userIdHeader, 10, 64)
		if err != nil {
			checkErroType(w, helpers.INTERNAL, err)
			return
		}
		dto.UserID = int(userId)
		typeError, err := hd.portfolioService.Add(dto)
		if err != nil {
			checkErroType(w, typeError, err)
			return
		}
	}
}
func (hd *HandlerService) GetInvestmentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdHeader := r.Header.Get("user_id")
		userId, err := strconv.ParseInt(userIdHeader, 10, 64)
		if err != nil {
			checkErroType(w, helpers.INTERNAL, err)
			return
		}
		serviceresult, typeError, err := hd.portfolioService.GetAll(int(userId))
		if err != nil {
			if err != nil {
				checkErroType(w, typeError, err)
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
			checkErroType(w, helpers.VALIDATION, err)
			return
		}
		serviceresult, typeError, err := hd.portfolioService.GetById(uint(id))
		if err != nil {
			checkErroType(w, typeError, err)
			return

		}

		response.JSONResponse(w, http.StatusOK, serviceresult)
	}
}
func (hd *HandlerService) UpdateInvestmentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseInt(params["idTransaction"], 10, 64)
		if err != nil {
			checkErroType(w, helpers.VALIDATION, err)
			return
		}
		var dto dtos.Portfolio
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&dto); err != nil {
			checkErroType(w, helpers.INTERNAL, err)
			return
		}
		userIdHeader := r.Header.Get("user_id")
		userId, err := strconv.ParseInt(userIdHeader, 10, 64)
		if err != nil {
			response.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}
		dto.Id = uint(id)
		dto.UserID = int(userId)
		typeError, err := hd.portfolioService.Update(dto)
		if err != nil {
			checkErroType(w, typeError, err)
			return
		}

	}
}
func (hd *HandlerService) DeleteInvestmentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseInt(params["idTransaction"], 10, 64)
		if err != nil {
			checkErroType(w, helpers.VALIDATION, err)
			return
		}
		userIdHeader := r.Header.Get("user_id")
		userId, err := strconv.ParseInt(userIdHeader, 10, 64)
		if err != nil {
			checkErroType(w, helpers.INTERNAL, err)
			return
		}
		typeError, err := hd.portfolioService.Delete(id, userId)
		if err != nil {
			checkErroType(w, typeError, err)
			return
		}
	}
}

func checkErroType(w http.ResponseWriter, typeError helpers.ErrorType, err error) {

	switch typeError {
	case helpers.VALIDATION:
		response.ErrorResponse(w, http.StatusBadRequest, err)
		return
	case helpers.MISSING:
		response.ErrorResponse(w, http.StatusNotFound, err)
		return
	case helpers.STATE:
		response.ErrorResponse(w, http.StatusConflict, err)
		return
	default:
		response.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

}
