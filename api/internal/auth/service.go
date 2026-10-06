package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/GeuberLucas/Gofre/api/internal/auth/security"
	"github.com/GeuberLucas/Gofre/api/internal/profile"
	dtos "github.com/GeuberLucas/Gofre/api/pkg/DTOs"
	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
	"github.com/GeuberLucas/Gofre/api/pkg/types"
)

type IAuthService interface {
	Login(obj dtos.LoginDTO) (*dtos.LoginResultDto, helpers.ErrorType, error)
	Register(obj dtos.RegisterDTO) (*dtos.LoginResultDto, helpers.ErrorType, error)

	ForgotPassword(email string) (helpers.ErrorType, error)
	ResetPassword(token string, newPassword string) (helpers.ErrorType, error)
}
type EmailMessage struct {
	TokenReset string `json:"tokenReset"`
	EmailTo    string `json:"emailTo"`
}
type AuthService struct {
	repository     IAuhtRepository
	profileService profile.IProfileService
}

func NewAuthService(repo IAuhtRepository, profileSvc profile.IProfileService) *AuthService {
	return &AuthService{
		repository:     repo,
		profileService: profileSvc,
	}
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

// Register implements [IAuthService].
func (s *AuthService) Register(obj dtos.RegisterDTO) (*dtos.LoginResultDto, helpers.ErrorType, error) {
	var usuario User

	passwordHash, err := security.HashPassword(obj.Password)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}

	// Agora apenas mapeamos o essencial
	usuario.Email = obj.Email
	usuario.Password = passwordHash
	usuario.Username = obj.Username

	if !usuario.Validate() {
		return nil, helpers.VALIDATION, errors.New("Required fields are empty")
	}

	// 1. Cria na tabela auth.users
	id, err := s.repository.CreateUser(usuario)
	if err != nil {
		return nil, helpers.INTERNAL, fmt.Errorf("User Not created: %s", err)
	}

	// 2. Cria o Perfil base com dados provisórios
	newProfile := profile.Profile{
		UserId:         id,
		CompleteName:   obj.Username, // Usamos o username provisoriamente pois o full_name não aceita null no DB
		CellPhone:      "",
		InitialBalance: types.FloatToMoney(0.0),
	}

	// Chama o serviço do módulo de profile para guardar o perfil
	_, err = s.profileService.AddProfileService(newProfile)
	if err != nil {
		return nil, helpers.INTERNAL, fmt.Errorf("User created but Profile base failed: %s", err)
	}

	// 3. Gera o Token e conclui o Login automático
	jwtToken, _ := security.GenerateToken(id)
	var result dtos.LoginResultDto
	result.Token = jwtToken

	return &result, helpers.NONE, nil
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
