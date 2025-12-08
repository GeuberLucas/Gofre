package dtos

import (
	"time"

	"github.com/GeuberLucas/Gofre/backend/pkg/types"
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/models"
)

type ExpenseDto struct {
	ID            int64     `json:"id"`
	UserId        int64     `json:"userId"`
	Description   string    `json:"description"`
	Target        string    `json:"target"`
	Category      string    `json:"category"`
	Type          string    `json:"type"`
	PaymentMethod string    `json:"paymentMethod"`
	PaymentDate   time.Time `json:"paymentDate"`
	IsPaid        bool      `json:"isPaid"`
	Amount        float64   `json:"amount"`
}

func (ex ExpenseDto) ToModel() models.Expense {
	amount := types.FloatToMoney(ex.Amount)
	return models.Expense{
		ID:            ex.ID,
		UserId:        ex.UserId,
		Description:   ex.Description,
		Target:        ex.Target,
		Category:      ex.Category,
		Type:          ex.Type,
		PaymentMethod: ex.PaymentMethod,
		PaymentDate:   ex.PaymentDate,
		IsPaid:        ex.IsPaid,
		Amount:        amount,
	}
}

type PaymentMethod string

const (
	PaymentMethodPix      PaymentMethod = "pix"
	PaymentMethodDebito   PaymentMethod = "debito"
	PaymentMethodCredito  PaymentMethod = "credito"
	PaymentMethodBoleto   PaymentMethod = "boleto"
	PaymentMethodDinheiro PaymentMethod = "dinheiro"
	PaymentMethodTED      PaymentMethod = "ted"
	PaymentMethodCheque   PaymentMethod = "cheque"
)
