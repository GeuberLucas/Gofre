package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/GeuberLucas/Gofre/backend/pkg/db"
	jwtToken "github.com/GeuberLucas/Gofre/backend/pkg/jwt"
	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
	dtos "github.com/GeuberLucas/Gofre/backend/services/auth/internal/DTOs"
	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/models"
	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/repository"
	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/security"
)

type EmailMessage struct {
	TokenReset string `json:"tokenReset"`
	EmailTo    string `json:"emailTo"`
}
type authService struct{}

func NewAuthService() *authService {
	return &authService{}
}

func (s *authService) Login(obj dtos.LoginDTO) (*dtos.LoginResultDto, error, string) {
	dbConn, userRepository, err := getUserRepository()
	defer dbConn.Close()
	if err != nil {
		return nil, err, "Internal"
	}

	userModel, err := userRepository.GetUserByUsername(obj.Username)

	if err != nil {
		return nil, err, "Internal"
	}
	passwordIsChecked := security.CheckPasswordHash(obj.Password, userModel.Password)

	if !passwordIsChecked {
		return nil, errors.New("Username or Password Invalids"), "Pass"
	}

	jwtToken, _ := jwtToken.GenerateToken(int(userModel.ID))

	var result dtos.LoginResultDto

	result.Token = jwtToken

	return &result, nil, ""
}

func (s *authService) Register(obj dtos.RegisterDTO) (*dtos.LoginResultDto, error, string) {
	var nameSplit []string = strings.SplitAfterN(obj.CompleteName, " ", 2)
	var usuario models.User
	passwordHash, err := security.HashPassword(obj.Password)
	if err != nil {
		return nil, err, "Internal"
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
		return nil, errors.New("Required fields are empty"), "validation"
	}
	dbConn, repositoryUser, err := getUserRepository()
	defer dbConn.Close()
	if err != nil {
		return nil, err, "Internal"
	}
	id := repositoryUser.CreateUser(usuario)
	if err != nil {
		return nil, fmt.Errorf("User Not created: %s", err), "Internal"
	}
	jwtToken, _ := jwtToken.GenerateToken(int(id))

	var result dtos.LoginResultDto

	result.Token = jwtToken

	return &result, nil, ""

}
func (s *authService) Profile(userID int64) (*dtos.ProfileDto, error, string) {

	dbConn, repositoryUser, err := getUserRepository()
	defer dbConn.Close()
	userModel, err := repositoryUser.GetUserByID(userID)
	if err != nil {
		return nil, err, "Internal"
	}
	var profileDto dtos.ProfileDto
	profileDto.CellPhone = userModel.Cellphone
	profileDto.Email = userModel.Email
	profileDto.FirstName = userModel.Name
	profileDto.LastName = userModel.LastName
	profileDto.UserID = userModel.ID

	return &profileDto, nil, ""
}

func (s *authService) ForgotPassword(email string) error {
	dbConn, userRepository, err := getUserRepository()
	defer dbConn.Close()
	if err != nil {
		return err
	}
	dbConn, resetTokenRepository, err := getResetTokenRepository()
	if err != nil {
		return err
	}
	user, err := userRepository.GetUserByEmail(email)
	if err != nil {
		return err
	}

	token, hashToken, err := security.CreateResetToken(32)
	if err != nil {
		return err
	}

	var resetTokenModel models.ResetToken
	resetTokenModel.UserID = user.ID
	resetTokenModel.TokenHash = hashToken.TokenHash
	resetTokenModel.ExpiresAt = hashToken.ExpiresAt

	err = resetTokenRepository.CreateResetToken(&resetTokenModel)
	if err != nil {
		return err
	}

	sendEmail(token, user.Email)

	return nil
}

func (s *authService) ResetPassword(token string, newPassword string) error {
	dbConn, userRepository, err := getUserRepository()
	dbConn, resetTokenRepository, err := getResetTokenRepository()
	defer dbConn.Close()
	if err != nil {
		return err
	}
	hashRecievedToken := security.HashToken(token)

	resetTokenModel, err := resetTokenRepository.GetResetTokenByTokenHash(hashRecievedToken)
	if err != nil {
		return err
	}
	if resetTokenModel.ExpiresAt.Unix() < time.Now().Unix() {
		return errors.New("Expired")
	}
	hashNewPassWord, err := security.HashPassword(newPassword)
	user, err := userRepository.GetUserByID(resetTokenModel.UserID)
	if err != nil {
		return err
	}
	userRepository.UpdateUserPassword(user.ID, hashNewPassWord)
	return nil
}

// TODO:Criar service de envio de email com comunicação por fila e pub,sub
func sendEmail(token string, email string) {
	var emailObj EmailMessage
	emailObj.TokenReset = token
	emailObj.EmailTo = email
	emailData, err := json.Marshal(emailObj)
	if err != nil {
		return
	}
	messaging.PublishMessage("auth.comunication.forgotPassword", emailData)
}

func getUserRepository() (*sql.DB, *repository.UserRepository, error) {
	dbConn, err := db.ConnectToDatabase()
	if err != nil {
		return dbConn, nil, err
	}
	return dbConn, repository.NewUserRepository(dbConn), nil
}
func getResetTokenRepository() (*sql.DB, *repository.ResetTokensRepository, error) {
	dbConn, err := db.ConnectToDatabase()
	if err != nil {
		return dbConn, nil, err
	}
	return dbConn, repository.NewResetTokensRepository(dbConn), nil
}
