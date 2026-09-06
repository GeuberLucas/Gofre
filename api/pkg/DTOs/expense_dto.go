package dtos

import (
	"time"

	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
)

type ExpenseDto struct {
	ID            uint                    `json:"id"`
	UserId        uint                    `json:"userId"`
	Description   string                  `json:"description"`
	Target        string                  `json:"target"`
	Category      helpers.ExpenseCategory `json:"category"`
	Type          helpers.ExpenseType     `json:"type"`
	PaymentMethod helpers.PaymentMethod   `json:"paymentMethod"`
	PaymentDate   time.Time               `json:"paymentDate"`
	IsPaid        bool                    `json:"isPaid"`
	Amount        float64                 `json:"amount"`
}
