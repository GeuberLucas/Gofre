package investments

import (
	dtos "github.com/GeuberLucas/Gofre/api/pkg/DTOs"
	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
	"github.com/GeuberLucas/Gofre/api/pkg/types"
)

type IPortfolioService interface {
	Add(dto dtos.Portfolio) (helpers.ErrorType, error)
	GetAll(userId int) ([]dtos.Portfolio, helpers.ErrorType, error)
	GetById(id uint) (dtos.Portfolio, helpers.ErrorType, error)
	Update(dto dtos.Portfolio) (helpers.ErrorType, error)
	Delete(id int64, userId int64) (helpers.ErrorType, error)
	UpdateIsDone(id uint, isDone bool) (helpers.ErrorType, error)
}

type PortfolioService struct {
	portifolioRepository IPortfolioRepository
}

// Add implements IPortfolioService.
func (p *PortfolioService) Add(dto dtos.Portfolio) (helpers.ErrorType, error) {
	portfolioModel := p.MapperDtoToModel(dto)
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
	_, err := p.portifolioRepository.GetById(uint(id))
	if err != nil {
		return helpers.INTERNAL, err
	}
	err = p.portifolioRepository.Delete(id, userId)
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
		portfolioDto := p.MapperModelToDto(portfolioModel)
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

	portfolioDto := p.MapperModelToDto(portfolioModel)

	return portfolioDto, helpers.NONE, nil

}

// Update implements IPortfolioService.
func (p *PortfolioService) Update(dto dtos.Portfolio) (helpers.ErrorType, error) {

	portfolioModel := p.MapperDtoToModel(dto)
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
func (p *PortfolioService) UpdateIsDone(id uint, isDone bool) (helpers.ErrorType, error) {

	portfolioModel, err := p.portifolioRepository.GetById(id)
	portfolioModel.IsDone = isDone

	if err != nil {
		return helpers.INTERNAL, err
	}

	err = p.portifolioRepository.Update(portfolioModel)
	if err != nil {
		return helpers.INTERNAL, err
	}

	return helpers.NONE, nil
}

func NewPortfolioService(repo IPortfolioRepository) IPortfolioService {
	return &PortfolioService{
		portifolioRepository: repo,
	}
}

func (p *PortfolioService) MapperModelToDto(model Portfolio) dtos.Portfolio {
	return dtos.Portfolio{
		Id:          model.Id,
		UserID:      model.User_id,
		Description: model.Description,
		Broker:      model.Broker,
		IsDone:      model.IsDone,
		DepositDate: model.Deposit_date,
		AssetID:     model.Asset_id,
		Amount:      model.Amount.ToFloat(),
	}
}

func (p *PortfolioService) MapperDtoToModel(dto dtos.Portfolio) Portfolio {
	return Portfolio{
		Id:           dto.Id,
		User_id:      dto.UserID,
		Description:  dto.Description,
		Broker:       dto.Broker,
		IsDone:       dto.IsDone,
		Deposit_date: dto.DepositDate,
		Asset_id:     dto.AssetID,
		Amount:       types.FloatToMoney(dto.Amount),
	}
}
