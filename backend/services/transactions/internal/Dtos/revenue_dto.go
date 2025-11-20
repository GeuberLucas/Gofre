package dtos

import (
	"time"

	"github.com/GeuberLucas/Gofre/backend/pkg/types"
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/models"
)

type RevenueDto struct {
	ID          int64     `json:"id"`
	UserId      int64     `json:"userId"`
	Description string    `json:"description"`
	Origin      string    `json:"origin"`
	Type        string    `json:"type"`
	ReceiveDate time.Time `json:"receiveDate"`
	IsRecieved  bool      `json:"IsRecieved"`
	Amount      float64   `json:"amount"`
}

func (re RevenueDto) ToModel() models.Revenue {
	return models.Revenue{
		ID:          re.ID,
		UserId:      re.UserId,
		Description: re.Description,
		Origin:      re.Origin,
		Type:        re.Type,
		ReceiveDate: re.ReceiveDate,
		IsRecieved:  re.IsRecieved,
		Amount:      types.FloatToMoney(re.Amount),
	}
}
