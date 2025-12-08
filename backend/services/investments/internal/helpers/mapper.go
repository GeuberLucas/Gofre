package helpers

import (
	"github.com/GeuberLucas/Gofre/backend/pkg/types"
	dtos "github.com/GeuberLucas/Gofre/backend/services/investments/internal/DTOs"
	"github.com/GeuberLucas/Gofre/backend/services/investments/internal/models"
)

func MapperDtoToModel(dto dtos.Portfolio) models.Portfolio {
	return models.Portfolio{
		Id:           dto.Id,
		User_id:      dto.UserID,
		Asset_id:     dto.AssetID,
		Deposit_date: dto.DepositDate,
		Broker:       dto.Broker,
		Amount:       types.FloatToMoney(dto.Amount),
		IsDone:       dto.IsDone,
		Description:  dto.Description,
	}
}

func MapperModelToDto(model models.Portfolio) dtos.Portfolio {
	return dtos.Portfolio{
		Id:          model.Id,
		UserID:      model.User_id,
		AssetID:     model.Asset_id,
		DepositDate: model.Deposit_date,
		Broker:      model.Broker,
		Amount:      model.Amount.ToFloat(),
		IsDone:      model.IsDone,
		Description: model.Description,
	}
}
