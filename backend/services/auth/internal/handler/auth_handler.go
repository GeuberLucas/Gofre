package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/GeuberLucas/Gofre/backend/pkg"
	dtos "github.com/GeuberLucas/Gofre/backend/services/auth/internal/DTOs"
	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/service"
	"github.com/gorilla/mux"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	bodyRequest,erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		pkg.ErrorResponse(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var loginDTO dtos.LoginDTO
	if erro = json.Unmarshal(bodyRequest, &loginDTO);erro != nil {
		pkg.ErrorResponse(w, http.StatusBadRequest, erro)
		return
	}
	service := service.NewAuthService()
 	serviceresult,erro,typeError := service.Login(loginDTO)
	if erro != nil {
		if typeError == "validation" {
			pkg.ErrorResponse(w, http.StatusBadRequest, erro)
			return
		}
		if typeError == "Internal" {
			pkg.ErrorResponse(w, http.StatusInternalServerError, erro)
			return
		}
	}

	pkg.JSONResponse(w,http.StatusOK,serviceresult)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	bodyRequest,erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		pkg.ErrorResponse(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var registerDTO dtos.RegisterDTO
	if erro = json.Unmarshal(bodyRequest, &registerDTO);erro != nil {
		pkg.ErrorResponse(w, http.StatusBadRequest, erro)
		return
	}
	service := service.NewAuthService()
	serviceresult,erro,typeError := service.Register(registerDTO)
	if erro != nil {
		if typeError == "validation" {
			pkg.ErrorResponse(w, http.StatusBadRequest, erro)
			return
		}
		if typeError == "Internal" {
			pkg.ErrorResponse(w, http.StatusInternalServerError, erro)
			return
		}
	}

	pkg.JSONResponse(w,http.StatusOK,serviceresult)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	params:= mux.Vars(r)
	userId,erro :=strconv.ParseUint(params["userId"],10,64)
	if erro != nil {
		pkg.ErrorResponse(w, http.StatusBadRequest, erro)
		return
	}
	service := service.NewAuthService()
	serviceresult,erro,typeError := service.Profile(strconv.FormatUint(userId,10))
	if erro != nil {
		if typeError == "validation" {
			pkg.ErrorResponse(w, http.StatusBadRequest, erro)
			return
		}
		if typeError == "Internal" {
			pkg.ErrorResponse(w, http.StatusInternalServerError, erro)
			return
		}
	}

	pkg.JSONResponse(w,http.StatusOK,serviceresult)
}

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	bodyRequest,erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		pkg.ErrorResponse(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var forgotPasswordDTO dtos.ForgotPasswordDTO
	if erro = json.Unmarshal(bodyRequest, &forgotPasswordDTO);erro != nil {
		pkg.ErrorResponse(w, http.StatusBadRequest, erro)
		return
	}
	service := service.NewAuthService()
	erro = service.ForgotPassword(forgotPasswordDTO.Email)
	if erro != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}
	pkg.JSONResponse(w,http.StatusOK,map[string]string{"message":"If the email exists, a reset link has been sent."})
}

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	params:= mux.Vars(r)
	HashEncoded :=params["HashEncoded"]
	if HashEncoded == "" {
		pkg.ErrorResponse(w, http.StatusBadRequest, nil)
		return
	}
	service := service.NewAuthService()
	bodyRequest,erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		pkg.ErrorResponse(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var resetPasswordDTO dtos.ResetPasswordDTO
	if erro = json.Unmarshal(bodyRequest, &resetPasswordDTO);erro != nil {
		pkg.ErrorResponse(w, http.StatusBadRequest, erro)
		return
	}
	erro = service.ResetPassword(HashEncoded,resetPasswordDTO.NewPassword)
	if erro != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}
	pkg.JSONResponse(w,http.StatusOK,map[string]string{"message":"Password has been reset successfully."})

}