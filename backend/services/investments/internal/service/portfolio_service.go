package service

import (
	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
	"github.com/GeuberLucas/Gofre/backend/pkg/types"
	dtos "github.com/GeuberLucas/Gofre/backend/services/investments/internal/DTOs"
	"github.com/GeuberLucas/Gofre/backend/services/investments/internal/repository"
)

type IPortfolioService interface {
	Add(dto dtos.Portfolio) error
	GetAll(userId int) ([]dtos.Portfolio, types.ErrorType, error)
	GetById(id uint) (dtos.Portfolio, error)
}

type PortfolioService struct {
	portifolioRepository repository.IPortfolioRepository
	broker               messaging.IMessaging
}

func NewPortfolioService(repo repository.IPortfolioRepository, broker messaging.IMessaging) IPortfolioService {
	return &PortfolioService{
		portifolioRepository: repo,
		broker:               broker,
	}
}
