package revenue

import (
	dtos "github.com/GeuberLucas/Gofre/api/pkg/DTOs"
	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
)

type IRevenueService interface {
	Add(dto dtos.RevenueDto) (helpers.ErrorType, error)
	GetAll(userId uint) ([]dtos.RevenueDto, helpers.ErrorType, error)
	GetById(id uint) (dtos.RevenueDto, helpers.ErrorType, error)
	Update(id uint, dto dtos.RevenueDto) (helpers.ErrorType, error)
	Delete(id uint, userId uint) (helpers.ErrorType, error)
	UpdateIsReceived(id uint, isDone bool) (helpers.ErrorType, error)
}
type RevenueService struct {
	revenueRepository IRevenueRepository
}

func NewRevenueService(repository IRevenueRepository) IRevenueService {
	return &RevenueService{
		revenueRepository: repository,
	}
}
func (ts *RevenueService) Add(dto dtos.RevenueDto) (helpers.ErrorType, error) {
	model := ToModel(dto)
	err := model.Isvalid()
	if err != nil {
		return helpers.VALIDATION, err
	}

	err = ts.revenueRepository.Create(model)
	if err != nil {
		return helpers.INTERNAL, err
	}

	if err != nil {
		return helpers.INTERNAL, err
	}
	return helpers.NONE, nil
}
func (ts *RevenueService) GetById(id uint) (dtos.RevenueDto, helpers.ErrorType, error) {
	revenueModel, err := ts.revenueRepository.GetById(id)
	if err != nil {
		return dtos.RevenueDto{}, helpers.INTERNAL, err
	}
	revenueDto := revenueDtoFromModel(revenueModel)
	return revenueDto, helpers.NONE, nil
}

func (ts *RevenueService) GetAll(idUser uint) ([]dtos.RevenueDto, helpers.ErrorType, error) {

	revenueModels, err := ts.revenueRepository.GetAll(idUser)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}
	var revenues []dtos.RevenueDto
	for _, revenueModel := range revenueModels {

		revenueDto := revenueDtoFromModel(revenueModel)
		revenues = append(revenues, revenueDto)
	}
	return revenues, helpers.NONE, nil
}
func (ts *RevenueService) UpdateIsReceived(id uint, isReceived bool) (helpers.ErrorType, error) {
	oldModel, err := ts.revenueRepository.GetById(id)
	if err != nil {
		return helpers.INTERNAL, err
	}
	model := oldModel
	model.IsRecieved = isReceived

	err = model.Isvalid()
	if err != nil {
		return helpers.VALIDATION, err
	}
	err = ts.revenueRepository.Update(model)
	if err != nil {
		return helpers.INTERNAL, err
	}

	if err != nil {
		return helpers.INTERNAL, err
	}
	return helpers.NONE, nil
}
func (ts *RevenueService) Update(id uint, dto dtos.RevenueDto) (helpers.ErrorType, error) {

	model := ToModel(dto)
	model.ID = id
	err := model.Isvalid()
	if err != nil {
		return helpers.VALIDATION, err
	}
	err = ts.revenueRepository.Update(model)
	if err != nil {
		return helpers.INTERNAL, err
	}

	return helpers.NONE, nil
}

func (ts *RevenueService) Delete(id uint, userId uint) (helpers.ErrorType, error) {
	_, err := ts.revenueRepository.GetById(id)
	if err != nil {
		return helpers.INTERNAL, err
	}
	err = ts.revenueRepository.Delete(id, userId)
	if err != nil {
		return helpers.INTERNAL, err
	}

	return helpers.NONE, nil
}
func revenueDtoFromModel(re Revenue) dtos.RevenueDto {
	return dtos.RevenueDto{
		ID:          int64(re.ID),
		UserId:      int64(re.UserId),
		Description: re.Description,
		Origin:      re.Origin,
		Type:        re.Type,
		ReceiveDate: re.ReceiveDate,
		IsRecieved:  re.IsRecieved,
		Amount:      re.Amount.ToFloat(),
	}
}
