package models

import (
	"errors"
	"time"

	"github.com/GeuberLucas/Gofre/backend/pkg/types"
)

// portfolio represents a investment
type Portfolio struct {
	Id           uint
	User_id      int
	Asset_id     uint
	Deposit_date time.Time
	Broker       string
	Amount       types.Money
	IsDone       bool
	Description  string
}

func (p *Portfolio) IsValid() error {
	if p.Deposit_date.IsZero() {
		return errors.New("Portfolio struct: Validate: deposit date is required")
	}
	if p.Asset_id == 0 {
		return errors.New("Portfolio struct: Validate: asset type is required")
	}

	return nil
}
