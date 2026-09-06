package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/GeuberLucas/Gofre/api/internal/auth/security"
	dtos "github.com/GeuberLucas/Gofre/api/pkg/DTOs"
	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
)

type IAuthService interface {
	Login(obj dtos.LoginDTO) (*dtos.LoginResultDto, helpers.ErrorType, error)
	Register(obj dtos.RegisterDTO) (*dtos.LoginResultDto, helpers.ErrorType, error)
	Profile(userID uint) (*dtos.ProfileDto, helpers.ErrorType, error)
	ForgotPassword(email string) (helpers.ErrorType, error)
	ResetPassword(token string, newPassword string) (helpers.ErrorType, error)
}
type EmailMessage struct {
	TokenReset string `json:"tokenReset"`
	EmailTo    string `json:"emailTo"`
}
type AuthService struct {
	repository IAuhtRepository
}

func NewAuthService(repo IAuhtRepository) *AuthService {
	return &AuthService{repository: repo}
}

func (s *AuthService) Login(obj dtos.LoginDTO) (*dtos.LoginResultDto, helpers.ErrorType, error) {

	userModel, err := s.repository.GetUserByUsername(obj.Username)

	if err != nil {
		return nil, helpers.INTERNAL, err
	}
	passwordIsChecked := security.CheckPasswordHash(obj.Password, userModel.Password)

	if !passwordIsChecked {
		return nil, helpers.VALIDATION, errors.New("Username or Password Invalids")
	}

	jwtToken, _ := security.GenerateToken(uint(userModel.ID))

	var result dtos.LoginResultDto

	result.Token = jwtToken

	return &result, helpers.NONE, nil
}

func (s *AuthService) Register(obj dtos.RegisterDTO) (*dtos.LoginResultDto, helpers.ErrorType, error) {
	var nameSplit []string = strings.SplitAfterN(obj.CompleteName, " ", 2)
	var usuario User
	passwordHash, err := security.HashPassword(obj.Password)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}
	usuario.Email = obj.Email
	usuario.Password = passwordHash
	usuario.Username = obj.Username
	usuario.Cellphone = obj.Cellphone
	usuario.Name = nameSplit[0]
	if len(nameSplit) > 1 {
		usuario.LastName = nameSplit[1]
	}
	if !usuario.Validate() {
		return nil, helpers.VALIDATION, errors.New("Required fields are empty")
	}

	id, err := s.repository.CreateUser(usuario)
	if err != nil {
		return nil, helpers.INTERNAL, fmt.Errorf("User Not created: %s", err)
	}
	jwtToken, _ := security.GenerateToken(id)

	var result dtos.LoginResultDto

	result.Token = jwtToken

	return &result, helpers.NONE, nil

}
func (s *AuthService) Profile(userID uint) (*dtos.ProfileDto, helpers.ErrorType, error) {

	userModel, err := s.repository.GetUserByID(userID)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}
	var profileDto dtos.ProfileDto
	profileDto.CellPhone = userModel.Cellphone
	profileDto.Email = userModel.Email
	profileDto.FirstName = userModel.Name
	profileDto.LastName = userModel.LastName
	profileDto.UserID = userModel.ID

	return &profileDto, helpers.NONE, nil
}

func (s *AuthService) ForgotPassword(email string) (helpers.ErrorType, error) {

	user, err := s.repository.GetUserByEmail(email)
	if err != nil {
		return helpers.INTERNAL, err
	}

	token, hashToken, err := security.CreateResetToken(32)
	if err != nil {
		return helpers.INTERNAL, err
	}

	var resetTokenModel ResetToken
	resetTokenModel.UserID = user.ID
	resetTokenModel.TokenHash = hashToken.TokenHash
	resetTokenModel.ExpiresAt = hashToken.ExpiresAt

	err = s.repository.CreateResetToken(&resetTokenModel)
	if err != nil {
		return helpers.INTERNAL, err
	}

	s.sendEmail(token, user.Email)

	return helpers.NONE, nil
}

func (s *AuthService) ResetPassword(token string, newPassword string) (helpers.ErrorType, error) {

	hashRecievedToken := security.HashToken(token)

	resetTokenModel, err := s.repository.GetResetTokenByTokenHash(hashRecievedToken)
	if err != nil {
		return helpers.INTERNAL, err
	}
	if resetTokenModel.ExpiresAt.Unix() < time.Now().Unix() {
		return helpers.STATE, errors.New("Expired")
	}
	hashNewPassWord, err := security.HashPassword(newPassword)
	user, err := s.repository.GetUserByID(uint(resetTokenModel.UserID))
	if err != nil {
		return helpers.INTERNAL, err
	}
	s.repository.UpdateUserPassword(uint(user.ID), hashNewPassWord)
	return helpers.NONE, nil
}

// TODO:Criar service de envio de email com comunicação por fila e pub,sub
func (s *AuthService) sendEmail(token string, email string) {
	var emailObj EmailMessage
	emailObj.TokenReset = token
	emailObj.EmailTo = email
	_, err := json.Marshal(emailObj)
	if err != nil {
		return
	}
	//s.messasingBroker.PublishMessage("auth.comunication.forgotPassword", emailData)
}
