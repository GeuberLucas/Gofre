package dtos

import (
	"time"
)

type Portfolio struct {
	Id          uint      `json:"id"`
	UserID      int       `json:"user_id"`
	AssetID     uint      `json:"asset_id"`
	DepositDate time.Time `json:"deposit_date"`
	Broker      string    `json:"broker"`
	Amount      float64   `json:"amount"`
	IsDone      bool      `json:"is_done"`
	Description string    `json:"description"`
}
