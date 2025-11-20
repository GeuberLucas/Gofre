package models

import (
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

func (ex Expense) Isvalid() (bool, string) {

	if ex.UserId == 0 {
		return false, "expense:validate:UserId required"
	}
	if ex.Target == "" {
		return false, "expense:validate:Target required"
	}
	if ex.Category == "" {
		return false, "expense:validate:Category required"
	}

	if ex.Type == "" {
		return false, "expense:validate:Category required"
	}

	if ex.PaymentMethod == "" {
		return false, "expense:validate:PaymentMethod required"
	}

	if ex.PaymentDate.IsZero() {
		return false, "expense:validate:PaymentDate required"
	}
	if ex.Amount <= 0 {
		return false, "repense:validate:Amount not be equal or minor than zero"
	}
	return true, ""
}
