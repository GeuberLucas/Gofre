package service

import (
	"encoding/json"
	"fmt"

	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
	dtos "github.com/GeuberLucas/Gofre/backend/services/investments/internal/DTOs"
	"github.com/GeuberLucas/Gofre/backend/services/investments/internal/helpers"
	"github.com/GeuberLucas/Gofre/backend/services/investments/internal/models"
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
	err = p.sendMessagingToBroker(&portfolioModel, nil, messaging.ActionInsert)
	if err != nil {
		return helpers.INTERNAL, err
	}
	return helpers.NONE, nil

}

// Delete implements IPortfolioService.
func (p *PortfolioService) Delete(id int64, userId int64) (helpers.ErrorType, error) {
	portfolioModel, err := p.portifolioRepository.GetById(uint(id))
	if err != nil {
		return helpers.INTERNAL, err
	}
	err = p.portifolioRepository.Delete(id, userId)
	if err != nil {
		return helpers.INTERNAL, err
	}
	err = p.sendMessagingToBroker(&portfolioModel, nil, messaging.ActionDelete)

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
	oldPortfolioModel, err := p.portifolioRepository.GetById(dto.Id)
	if err != nil {
		return helpers.INTERNAL, err
	}
	portfolioModel := helpers.MapperDtoToModel(dto)
	err = portfolioModel.IsValid()
	if err != nil {
		return helpers.VALIDATION, err
	}

	err = p.portifolioRepository.Update(portfolioModel)
	if err != nil {
		return helpers.INTERNAL, err
	}
	err = p.sendMessagingToBroker(&portfolioModel, &oldPortfolioModel, messaging.ActionUpdate)
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

func (p *PortfolioService) sendMessagingToBroker(model *models.Portfolio, modelOld *models.Portfolio, action messaging.ActionType) error {

	ms := messaging.MessagingDto{
		Month:            model.Deposit_date.Month(),
		Year:             uint(model.Deposit_date.Local().Year()),
		Amount:           model.Amount,
		Movement:         messaging.TypeInvestment,
		MovementType:     models.GetAssetName(model.Asset_id),
		MovementCategory: "",
		WithCredit:       false,
		IsConfirmed:      model.IsDone,
		Action:           action,
		UserId:           model.User_id,
	}

	if action == messaging.ActionUpdate {
		ms.AmountOld = modelOld.Amount
		ms.MonthOld = modelOld.Deposit_date.Month()
		ms.YearOld = uint(modelOld.Deposit_date.Year())
		ms.MovementCategoryOld = ""
		ms.MovementTypeOld = models.GetAssetName(modelOld.Asset_id)
		ms.WithCreditOld = false
		ms.IsConfirmedOld = modelOld.IsDone
	}

	if err := ms.IsValid(); err != nil {
		return err
	}

	eventName := fmt.Sprintf("finance.%s.%s", messaging.TypeInvestment, action)
	json, err := json.Marshal(ms)
	if err != nil {
		return err
	}
	p.broker.PublishMessage(eventName, json)

	return nil

}
