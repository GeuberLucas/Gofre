package revenue

import (
	"github.com/!geuber!lucas/!gofre/backend/pkg/messaging"
	dtos "github.com/GeuberLucas/Gofre/api/pkg/DTOs"
	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
)

type IRevenueService interface {
	Add(dto dtos.Revenue) (helpers.ErrorType, error)
	GetAll(userId int) ([]dtos.Revenue, helpers.ErrorType, error)
	GetById(id uint) (dtos.Revenue, helpers.ErrorType, error)
	Update(dto dtos.Revenue) (helpers.ErrorType, error)
	Delete(id int64, userId int64) (helpers.ErrorType, error)
	UpdateIsDone(id uint, isDone bool) (helpers.ErrorType, error)
}

func (ts *TransactionService) AddRevenue(dto dtos.RevenueDto) (string, error) {
	model := dto.ToModel()
	err := model.Isvalid()
	if err != nil {
		return "Validation", err
	}

	err = ts.revenueRepository.Create(model)
	if err != nil {
		return "Internal", err
	}
	err = ts.sendRevenueToBroker(&model, nil, messaging.ActionInsert)
	if err != nil {
		return helpers.INTERNAL.String(), err
	}
	return "", err
}
func (ts *TransactionService) GetByIdRevenue(id int64) (dtos.RevenueDto, error, string) {
	revenueModel, err := ts.revenueRepository.GetById(id)
	if err != nil {
		return dtos.RevenueDto{}, err, "Internal"
	}
	revenueDto := revenueDtoFromModel(revenueModel)
	return revenueDto, nil, ""
}

func (ts *TransactionService) GetByIdUserRevenue(idUser int64) ([]dtos.RevenueDto, error, string) {

	revenueModels, err := ts.revenueRepository.GetByUserId(idUser)
	if err != nil {
		return nil, err, "Internal"
	}
	var revenues []dtos.RevenueDto
	for _, revenueModel := range revenueModels {

		revenueDto := revenueDtoFromModel(revenueModel)
		revenues = append(revenues, revenueDto)
	}
	return revenues, nil, ""
}
func (ts *TransactionService) UpdateIsReceivedRevenue(id int64, isReceived bool) (error, string) {
	oldModel, err := ts.revenueRepository.GetById(id)
	if err != nil {
		return err, "Internal"
	}
	model := oldModel
	model.IsRecieved = isReceived

	err = model.Isvalid()
	if err != nil {
		return err, "Validation"
	}
	err = ts.revenueRepository.Update(model)
	if err != nil {
		return err, "Internal"
	}
	err = ts.sendRevenueToBroker(&model, &oldModel, messaging.ActionUpdate)
	if err != nil {
		return err, helpers.INTERNAL.String()
	}
	return nil, ""
}
func (ts *TransactionService) UpdateRevenue(id int64, dto dtos.RevenueDto) (error, string) {
	oldModel, err := ts.revenueRepository.GetById(id)
	if err != nil {
		return err, "Internal"
	}
	model := dto.ToModel()
	model.ID = id
	err = model.Isvalid()
	if err != nil {
		return err, "Validation"
	}
	err = ts.revenueRepository.Update(model)
	if err != nil {
		return err, "Internal"
	}
	err = ts.sendRevenueToBroker(&model, &oldModel, messaging.ActionUpdate)
	if err != nil {
		return err, helpers.INTERNAL.String()
	}
	return nil, ""
}
func revenueDtoFromModel(re models.Revenue) dtos.RevenueDto {
	return dtos.RevenueDto{
		ID:          re.ID,
		UserId:      re.UserId,
		Description: re.Description,
		Origin:      re.Origin,
		Type:        re.Type,
		ReceiveDate: re.ReceiveDate,
		IsRecieved:  re.IsRecieved,
		Amount:      re.Amount.ToFloat(),
	}
}
func (ts *TransactionService) DeleteRevenue(id int64, userId int64) (error, string) {
	model, err := ts.revenueRepository.GetById(id)
	if err != nil {
		return err, "Internal"
	}
	err = ts.revenueRepository.Delete(id, userId)
	if err != nil {
		return err, "Internal"
	}

	err = ts.sendRevenueToBroker(&model, nil, messaging.ActionDelete)
	if err != nil {
		return err, helpers.INTERNAL.String()
	}

	return nil, ""
}
