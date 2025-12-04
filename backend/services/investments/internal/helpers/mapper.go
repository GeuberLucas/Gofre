package helpers

import (
	"github.com/GeuberLucas/Gofre/backend/pkg/types"
	dtos "github.com/GeuberLucas/Gofre/backend/services/investments/internal/DTOs"
	"github.com/GeuberLucas/Gofre/backend/services/investments/internal/models"
)

func MapperDtoToModel(dto dtos.Portfolio) models.Portfolio {
	return models.Portfolio{
		Id:           dto.Id,
		User_id:      dto.User_id,
		Asset_id:     dto.Asset_id,
		Deposit_date: dto.Deposit_date,
		Broker:       dto.Broker,
		Amount:       types.FloatToMoney(dto.Amount),
		IsDone:       dto.IsDone,
		Description:  dto.Description,
	}
}

func MapperModelToDto(model models.Portfolio) dtos.Portfolio {
	return dtos.Portfolio{
		Id:           model.Id,
		User_id:      model.User_id,
		Asset_id:     model.Asset_id,
		Deposit_date: model.Deposit_date,
		Broker:       model.Broker,
		Amount:       model.Amount.ToFloat(),
		IsDone:       model.IsDone,
		Description:  model.Description,
	}
}
