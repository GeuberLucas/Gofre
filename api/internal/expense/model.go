package expense

import (
	"errors"
	"time"

	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
	"github.com/GeuberLucas/Gofre/api/pkg/types"
)

type Expense struct {
	ID            uint
	UserId        uint
	Description   string
	Target        string
	Category      helpers.ExpenseCategory
	Type          helpers.ExpenseType
	PaymentMethod helpers.PaymentMethod
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
	if ex.Category < 0 || ex.Category > helpers.ExpenseCategoryOutros {
		return errors.New("expense:validate:Category invalid")
	}

	if ex.Type < 0 || ex.Type > helpers.ExpenseTypeFatura {
		return errors.New("expense:validate:Type invalid")
	}

	if ex.PaymentMethod < 0 || ex.PaymentMethod > helpers.PaymentMethodCheque {
		return errors.New("expense:validate:PaymentMethod invalid")
	}

	if ex.PaymentDate.IsZero() {
		return errors.New("expense:validate:PaymentDate required")
	}
	if ex.Amount <= 0 {
		return errors.New("expense:validate:Amount not be equal or minor than zero")
	}
	return nil
}
