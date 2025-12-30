package service

import (
	"testing"

	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
	"github.com/GeuberLucas/Gofre/backend/pkg/types"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/models"
)

func TestCalculateValuesExpense(t *testing.T) {
	tests := []struct {
		name          string
		dto           messaging.MessagingDto
		model         *models.Expense
		amount        types.Money
		expectedError bool
		validate      func(*testing.T, *models.Expense)
	}{
		{
			name: "Invoice movement without credit confirmed",
			dto: messaging.MessagingDto{
				MovementType: "Fatura",
				WithCredit:   false,
				IsConfirmed:  true,
			},
			model:         &models.Expense{},
			amount:        100,
			expectedError: false,
			validate: func(t *testing.T, m *models.Expense) {
				if m.Invoice != 100 || m.Planned != 100 || m.Actual != 100 {
					t.Errorf("expected Invoice=100, Planned=100, Actual=100; got Invoice=%v, Planned=%v, Actual=%v", m.Invoice, m.Planned, m.Actual)
				}
			},
		},
		{
			name: "Monthly movement without credit pending",
			dto: messaging.MessagingDto{
				MovementType: "Mensal",
				WithCredit:   false,
				IsConfirmed:  false,
			},
			model:         &models.Expense{},
			amount:        50,
			expectedError: false,
			validate: func(t *testing.T, m *models.Expense) {
				if m.Monthly != 50 || m.Planned != 50 || m.Pending != 50 {
					t.Errorf("expected Monthly=50, Planned=50, Pending=50; got Monthly=%v, Planned=%v, Pending=%v", m.Monthly, m.Planned, m.Pending)
				}
			},
		},
		{
			name: "Variable movement with credit",
			dto: messaging.MessagingDto{
				MovementType: "Variavel",
				WithCredit:   true,
				IsConfirmed:  true,
			},
			model:         &models.Expense{},
			amount:        75,
			expectedError: false,
			validate: func(t *testing.T, m *models.Expense) {
				if m.Variable != 75 || m.Planned != 0 || m.Actual != 0 {
					t.Errorf("expected Variable=75, Planned=0, Actual=0; got Variable=%v, Planned=%v, Actual=%v", m.Variable, m.Planned, m.Actual)
				}
			},
		},
		{
			name: "Invalid movement type",
			dto: messaging.MessagingDto{
				MovementType: "Invalid",
				WithCredit:   false,
				IsConfirmed:  true,
			},
			model:         &models.Expense{},
			amount:        100,
			expectedError: true,
			validate:      func(t *testing.T, m *models.Expense) {},
		},
		{
			name: "Negative amount deletion",
			dto: messaging.MessagingDto{
				MovementType: "Fatura",
				WithCredit:   false,
				IsConfirmed:  true,
			},
			model:         &models.Expense{Invoice: 100, Planned: 100, Actual: 100},
			amount:        -50,
			expectedError: false,
			validate: func(t *testing.T, m *models.Expense) {
				if m.Invoice != 50 || m.Planned != 50 || m.Actual != 50 {
					t.Errorf("expected Invoice=50, Planned=50, Actual=50; got Invoice=%v, Planned=%v, Actual=%v", m.Invoice, m.Planned, m.Actual)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := calculateValuesExpense(tt.dto, tt.model, tt.amount)
			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err != nil)
			}
			tt.validate(t, tt.model)
		})
	}
}
