package dtos

import (
	"time"

	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
)

type RevenueDto struct {
	ID          int64              `json:"id"`
	UserId      int64              `json:"userId"`
	Description string             `json:"description"`
	Origin      string             `json:"origin"`
	Type        helpers.IncomeType `json:"type"`
	ReceiveDate time.Time          `json:"receiveDate"`
	IsRecieved  bool               `json:"isRecieved"`
	Amount      float64            `json:"amount"`
}
