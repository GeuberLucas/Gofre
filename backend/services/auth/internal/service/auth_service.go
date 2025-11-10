package service

import dtos "github.com/GeuberLucas/Gofre/backend/services/auth/internal/DTOs"


type authService struct{}

func NewAuthService() *authService {
    return &authService{}
}

func (s *authService) Login(obj dtos.LoginDTO) (*dtos.LoginResultDto, error, string) {
	return nil, nil, ""
}