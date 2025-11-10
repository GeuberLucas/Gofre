package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/GeuberLucas/Gofre/backend/pkg"
	dtos "github.com/GeuberLucas/Gofre/backend/services/auth/internal/DTOs"
	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/service"
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
	w.Write([]byte("Register"))
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Profile"))
}

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Forgot Password"))
}

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Reset Password"))
}