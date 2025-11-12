package service

import (
	"errors"
	"strings"

	"github.com/GeuberLucas/Gofre/backend/pkg"
	dtos "github.com/GeuberLucas/Gofre/backend/services/auth/internal/DTOs"
	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/models"
	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/repository"
	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/security"
)


type authService struct{}

func NewAuthService() *authService {
    return &authService{}
}


func (s *authService) Login(obj dtos.LoginDTO) (*dtos.LoginResultDto, error, string) {
	userRepository,err:=getUserRepository()
	if err !=nil{
		return nil,err,"Internal"
	}
	
	userModel,err:= userRepository.GetUserByUsername(obj.Username)
	if err !=nil{
		return nil,err,"Internal"
	}
	passwordIsChecked:= security.CheckPasswordHash(obj.Password,userModel.Password)

	if !passwordIsChecked{
		return nil,errors.New("Username or Password Invalids"),"Pass"
	}

	jwtToken,_ := security.GenerateToken(int(userModel.ID))

	var result dtos.LoginResultDto

	result.Token=jwtToken
	

	return &result, nil, ""
}

func (s *authService) Register(obj dtos.RegisterDTO) (*dtos.LoginResultDto, error, string) {
	var nameSplit []string= strings.SplitAfterN(obj.CompleteName, " ", 2)
	var usuario models.User
	passwordHash,err:=security.HashPassword(obj.Password)
	if err != nil {
		return nil, err, "Internal"
	}
	usuario.Email = obj.Email
	usuario.Password = passwordHash
	usuario.Name = nameSplit[0]
	if len(nameSplit) > 1 {
		usuario.LastName = nameSplit[1]
	}

	repositoryUser,err := getUserRepository()
	if err != nil {
		return nil, err, "Internal"
	}
	id:= repositoryUser.CreateUser(usuario)
	if id <= 0 {
		return nil, errors.New("User not created"), "Internal"
	}
	jwtToken,_ := security.GenerateToken(int(id))

	var result dtos.LoginResultDto

	result.Token=jwtToken
	

	return &result, nil, ""


}
func (s *authService) Profile(userID int64) (*dtos.ProfileDto, error, string) {
	
	repositoryUser,err  := getUserRepository()
	userModel,err := repositoryUser.GetUserByID(userID)
	if err != nil{
		return nil,err,"Internal"
	}
	var profileDto dtos.ProfileDto 
	profileDto.CellPhone=userModel.Cellphone
	profileDto.Email=userModel.Email
	profileDto.FirstName=userModel.Name
	profileDto.LastName= userModel.LastName
	profileDto.UserID=userModel.ID
	
	return &profileDto, nil, ""
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



func getUserRepository() (*repository.UserRepository,error){
	dbConn,err := pkg.ConnectToDatabase()
	if err!= nil {
		return  nil,err
	}
	defer pkg.CloseDatabaseConnection(dbConn)
	return repository.NewUserRepository(dbConn),nil
}

