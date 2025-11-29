// Unit test
package models_test

import (
	"errors"
	"testing"
	"time"

	"github.com/GeuberLucas/Gofre/backend/pkg/types"
	. "github.com/GeuberLucas/Gofre/backend/services/transaction/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestValidationExpense(t *testing.T) {
	base := Expense{
		ID:            1,
		UserId:        1,
		Description:   "valid description",
		Target:        "valid",
		Category:      "valid",
		Type:          "valid",
		PaymentMethod: "valid",
		PaymentDate:   time.Date(2025, 10, 25, 0, 0, 0, 0, time.UTC),
		Amount:        types.FloatToMoney(50),
		IsPaid:        true,
	}

	type testCase struct {
		modify    func(e *Expense)
		wantValid bool
		wantErr   error
	}

	t.Run("test validation expense model", func(t *testing.T) {
		tests := []testCase{

			{modify: func(e *Expense) {}, wantValid: true, wantErr: nil},
			{modify: func(e *Expense) { e.UserId = 0 }, wantValid: false, wantErr: errors.New("expense:validate:UserId required")},
			{modify: func(e *Expense) { e.Target = "" }, wantValid: false, wantErr: errors.New("expense:validate:Target required")},
			{modify: func(e *Expense) { e.Category = "" }, wantValid: false, wantErr: errors.New("expense:validate:Category required")},
			{modify: func(e *Expense) { e.Type = "" }, wantValid: false, wantErr: errors.New("expense:validate:Type required")},
			{modify: func(e *Expense) { e.PaymentMethod = "" }, wantValid: false, wantErr: errors.New("expense:validate:PaymentMethod required")},
			{modify: func(e *Expense) { e.PaymentDate = time.Time{} }, wantValid: false, wantErr: errors.New("expense:validate:PaymentDate required")},
			{modify: func(e *Expense) { e.Amount = types.Money(-1) }, wantValid: false, wantErr: errors.New("expense:validate:Amount not be equal or minor than zero")},
			{modify: func(e *Expense) { e.Amount = types.Money(0) }, wantValid: false, wantErr: errors.New("expense:validate:Amount not be equal or minor than zero")},
		}

		for _, test := range tests {
			exp := base
			test.modify(&exp)
			gotErr := exp.Isvalid()

			if test.wantErr == nil {
				assert.NoError(t, gotErr)
			} else {
				assert.EqualError(t, gotErr, test.wantErr.Error())
			}

		}

	})
}
