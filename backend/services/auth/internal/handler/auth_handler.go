package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/GeuberLucas/Gofre/backend/pkg/response"
	dtos "github.com/GeuberLucas/Gofre/backend/services/auth/internal/DTOs"
	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/service"
	"github.com/gorilla/mux"
)

func LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	bodyRequest,erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.ErrorResponse(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var loginDTO dtos.LoginDTO
	if erro = json.Unmarshal(bodyRequest, &loginDTO);erro != nil {
		response.ErrorResponse(w, http.StatusBadRequest, erro)
		return
	}
	service := service.NewAuthService()
 	serviceresult,erro,typeError := service.Login(loginDTO)
	if erro != nil {
		if typeError == "validation" {
			response.ErrorResponse(w, http.StatusBadRequest, erro)
			return
		}
		if typeError == "Internal" {
			response.ErrorResponse(w, http.StatusInternalServerError, erro)
			return
		}
		if typeError == "Pass" {
			response.ErrorResponse(w, http.StatusUnauthorized, erro)
			return
		}
	}

	response.JSONResponse(w,http.StatusOK,serviceresult)
}}

func RegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	bodyRequest,erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.ErrorResponse(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var registerDTO dtos.RegisterDTO
	if erro = json.Unmarshal(bodyRequest, &registerDTO);erro != nil {
		response.ErrorResponse(w, http.StatusBadRequest, erro)
		return
	}
	service := service.NewAuthService()
	serviceresult,erro,typeError := service.Register(registerDTO)
	if erro != nil {
		if typeError == "validation" {
			response.ErrorResponse(w, http.StatusBadRequest, erro)
			return
		}
		if typeError == "Internal" {
			response.ErrorResponse(w, http.StatusInternalServerError, erro)
			return
		}
		
	}

	response.JSONResponse(w,http.StatusOK,serviceresult)
}}

func ProfileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	params:= mux.Vars(r)
	userId,erro :=strconv.ParseInt(params["userId"],10,64)
	if erro != nil {
		response.ErrorResponse(w, http.StatusBadRequest, erro)
		return
	}
	service := service.NewAuthService()
	serviceresult,erro,typeError := service.Profile(userId)
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

	response.JSONResponse(w,http.StatusOK,serviceresult)
}}

func ForgotPasswordHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	bodyRequest,erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.ErrorResponse(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var forgotPasswordDTO dtos.ForgotPasswordDTO
	if erro = json.Unmarshal(bodyRequest, &forgotPasswordDTO);erro != nil {
		response.ErrorResponse(w, http.StatusBadRequest, erro)
		return
	}
	service := service.NewAuthService()
	erro = service.ForgotPassword(forgotPasswordDTO.Email)
	if erro != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSONResponse(w,http.StatusOK,map[string]string{"message":"If the email exists, a reset link has been sent."})
}}

func ResetPasswordHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	params:= mux.Vars(r)
	HashEncoded :=params["HashEncoded"]
	if HashEncoded == "" {
		response.ErrorResponse(w, http.StatusBadRequest, nil)
		return
	}
	service := service.NewAuthService()
	bodyRequest,erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.ErrorResponse(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var resetPasswordDTO dtos.ResetPasswordDTO
	if erro = json.Unmarshal(bodyRequest, &resetPasswordDTO);erro != nil {
		response.ErrorResponse(w, http.StatusBadRequest, erro)
		return
	}
	erro = service.ResetPassword(HashEncoded,resetPasswordDTO.NewPassword)
	if erro != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSONResponse(w,http.StatusOK,map[string]string{"message":"Password has been reset successfully."})

}}