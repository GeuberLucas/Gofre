package service

import (
	"github.com/GeuberLucas/Gofre/backend/pkg"
	dtos "github.com/GeuberLucas/Gofre/backend/services/auth/internal/DTOs"
)


type authService struct{}

func NewAuthService() *authService {
    return &authService{}
}

func (s *authService) Login(obj dtos.LoginDTO) (*dtos.LoginResultDto, error, string) {
	dbConn,error := pkg.ConnectToDatabase()
	if error != nil {
		return nil, error, "Internal"
	}
	defer pkg.CloseDatabaseConnection(dbConn)
	return nil, nil, ""
}

func (s *authService) Register(obj dtos.RegisterDTO) (*dtos.LoginResultDto, error, string) {
	dbConn,err := pkg.ConnectToDatabase()
	if err != nil {
		return nil, err, "Internal"
	}
	defer pkg.CloseDatabaseConnection(dbConn)
	return nil, nil, ""
}
func (s *authService) Profile(userID string) (*dtos.ProfileDto, error, string) {
	dbConn,err := pkg.ConnectToDatabase()
	if err != nil {
		return nil, err, "Internal"
	}
	defer pkg.CloseDatabaseConnection(dbConn)
	return nil, nil, ""
}

func (s *authService) ForgotPassword(email string) error {
	dbConn,err := pkg.ConnectToDatabase()
	if err != nil {
		return  err
	}
	defer pkg.CloseDatabaseConnection(dbConn)
	return nil
}

func (s *authService) ResetPassword(token string, newPassword string) error {
	dbConn,err := pkg.ConnectToDatabase()
	if err!= nil {
		return  err
	}
	defer pkg.CloseDatabaseConnection(dbConn)
	return nil
}



