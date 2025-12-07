package service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
	"github.com/GeuberLucas/Gofre/backend/pkg/types"
	dtos "github.com/GeuberLucas/Gofre/backend/services/investments/internal/DTOs"
	"github.com/GeuberLucas/Gofre/backend/services/investments/internal/helpers"
	"github.com/GeuberLucas/Gofre/backend/services/investments/internal/repository"
)

type IPortfolioService interface {
	Add(dto dtos.Portfolio) (helpers.ErrorType, error)
	GetAll(userId int) ([]dtos.Portfolio, helpers.ErrorType, error)
	GetById(id uint) (dtos.Portfolio, helpers.ErrorType, error)
	Update(dto dtos.Portfolio) (helpers.ErrorType, error)
	Delete(id int64, userId int64) (helpers.ErrorType, error)
}

type PortfolioService struct {
	portifolioRepository repository.IPortfolioRepository
	broker               messaging.IMessaging
}

// Add implements IPortfolioService.
func (p *PortfolioService) Add(dto dtos.Portfolio) (helpers.ErrorType, error) {
	portfolioModel := helpers.MapperDtoToModel(dto)
	err := portfolioModel.IsValid()
	if err != nil {
		return helpers.VALIDATION, err
	}

	err = p.portifolioRepository.Create(portfolioModel)
	if err != nil {
		return helpers.INTERNAL, err
	}

	return helpers.NONE, nil

}

// Delete implements IPortfolioService.
func (p *PortfolioService) Delete(id int64, userId int64) (helpers.ErrorType, error) {
	err := p.portifolioRepository.Delete(id, userId)
	if err != nil {
		return helpers.INTERNAL, err
	}

	return helpers.NONE, nil
}

// GetAll implements IPortfolioService.
func (p *PortfolioService) GetAll(userId int) ([]dtos.Portfolio, helpers.ErrorType, error) {
	investments, err := p.portifolioRepository.GetAll(userId)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}

	var protfolioDtos []dtos.Portfolio
	for _, portfolioModel := range investments {
		portfolioDto := helpers.MapperModelToDto(portfolioModel)
		protfolioDtos = append(protfolioDtos, portfolioDto)
	}
	return protfolioDtos, helpers.NONE, nil
}

// GetById implements IPortfolioService.
func (p *PortfolioService) GetById(id uint) (dtos.Portfolio, helpers.ErrorType, error) {
	portfolioModel, err := p.portifolioRepository.GetById(id)
	if err != nil {
		return dtos.Portfolio{}, helpers.INTERNAL, err
	}

	portfolioDto := helpers.MapperModelToDto(portfolioModel)

	return portfolioDto, helpers.NONE, nil

}

// Update implements IPortfolioService.
func (p *PortfolioService) Update(dto dtos.Portfolio) (helpers.ErrorType, error) {
	portfolioModel := helpers.MapperDtoToModel(dto)
	err := portfolioModel.IsValid()
	if err != nil {
		return helpers.VALIDATION, err
	}

	err = p.portifolioRepository.Update(portfolioModel)
	if err != nil {
		return helpers.INTERNAL, err
	}

	return helpers.NONE, nil
}

func NewPortfolioService(repo repository.IPortfolioRepository, broker messaging.IMessaging) IPortfolioService {
	return &PortfolioService{
		portifolioRepository: repo,
		broker:               broker,
	}
}

func (p *PortfolioService) sendMessagingToBroker(month uint,
	year uint,
	amount types.Money,
	movementType string,
	isConfirmed bool,
	action messaging.ActionType) error {
	ms, err := messaging.NewMessagingDto(
		month,
		year,
		amount,
		messaging.TypeInvestment,
		movementType,
		"",
		false,
		isConfirmed,
		action,
	)

	if err != nil {
		return err
	}

	log.Println(ms)
	eventName := fmt.Sprintf("finance.%s.%d", messaging.TypeInvestment, action)
	json, err := json.Marshal(ms)
	if err != nil {
		return err
	}
	p.broker.PublishMessage(eventName, json)

	return nil

}
