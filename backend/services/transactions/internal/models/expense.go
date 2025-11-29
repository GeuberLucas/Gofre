package models

import (
	"errors"
	"time"

	"github.com/GeuberLucas/Gofre/backend/pkg/types"
)

type Expense struct {
	ID            int64
	UserId        int64
	Description   string
	Target        string
	Category      string
	Type          string
	PaymentMethod string
	PaymentDate   time.Time
	Amount        types.Money
	IsPaid        bool
}

func (ex *Expense) Isvalid() error {

	if ex.UserId == 0 {
		return errors.New("expense:validate:UserId required")
	}
	if ex.Target == "" {
		return errors.New("expense:validate:Target required")
	}
	if ex.Category == "" {
		return errors.New("expense:validate:Category required")
	}

	if ex.Type == "" {
		return errors.New("expense:validate:Type required")
	}

	if ex.PaymentMethod == "" {
		return errors.New("expense:validate:PaymentMethod required")
	}

	if ex.PaymentDate.IsZero() {
		return errors.New("expense:validate:PaymentDate required")
	}
	if ex.Amount <= 0 {
		return errors.New("expense:validate:Amount not be equal or minor than zero")
	}
	return nil
}
