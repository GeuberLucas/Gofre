package service

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/GeuberLucas/Gofre/backend/pkg/helpers"
	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
	"github.com/GeuberLucas/Gofre/backend/pkg/types"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/mocks"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	//PRIMEIRO: SETUP DO QUE A FUNÇÃO PRECISA
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error create mock database: %v", err)
	}
	defer mockDB.Close()
	mock.ExpectBegin()
	sqlTransac, err := mockDB.Begin()
	if err != nil {
		t.Fatalf("error create mock transaction: %v", err)
	}
	subscriberDTO := messaging.MessagingDto{
		Month:    time.April,
		Year:     2025,
		UserId:   1,
		Amount:   10,
		Movement: messaging.TypeIncome,
		Action:   messaging.ActionInsert,
	}
	//SERVICE E REPOSITORY
	repositoryMock := mocks.NewMockIAggregatedRepository(ctrl)
	service := NewService(repositoryMock)

	err = service.RegisterEvent(sqlTransac, subscriberDTO, service.HasBasicChanges)
	if err != nil {
		t.Error(err)
	}
}

func TestCalculateValuesModel(t *testing.T) {

	testCases := []struct {
		name       string
		dto        messaging.MessagingDto
		model      models.Aggregated
		multiplier int

		check func(t *testing.T, model models.Aggregated)
	}{
		{
			name: "add value in income",
			dto: messaging.MessagingDto{
				Amount:   10,
				Movement: messaging.TypeIncome,
			},
			model:      models.Aggregated{},
			multiplier: 1,
			check: func(t *testing.T, model models.Aggregated) {
				assert.Equal(t, types.Money(10), model.Revenue)
			},
		},
		{
			name: "add value in variable without credit",
			dto: messaging.MessagingDto{
				Amount:       10,
				Movement:     messaging.TypeExpense,
				MovementType: string(helpers.ExpenseTypeVariavel),
				WithCredit:   false,
			},
			model:      models.Aggregated{},
			multiplier: 1,
			check: func(t *testing.T, model models.Aggregated) {
				assert.Equal(t, types.Money(10), model.Expense)
				assert.Equal(t, types.Money(10), model.VariableWithoutCredit)

			},
		},
		{
			name: "add value in variable with credit",
			dto: messaging.MessagingDto{
				Amount:       10,
				Movement:     messaging.TypeExpense,
				MovementType: string(helpers.ExpenseTypeVariavel),
				WithCredit:   true,
			},
			model:      models.Aggregated{},
			multiplier: 1,
			check: func(t *testing.T, model models.Aggregated) {
				assert.Equal(t, types.Money(10), model.Expense)
				assert.Equal(t, types.Money(10), model.VariableWithCredit)

			},
		},
		{
			name: "add value in montlhy without credit",
			dto: messaging.MessagingDto{
				Amount:       10,
				Movement:     messaging.TypeExpense,
				MovementType: string(helpers.ExpenseTypeMensal),
				WithCredit:   false,
			},
			model:      models.Aggregated{},
			multiplier: 1,
			check: func(t *testing.T, model models.Aggregated) {
				assert.Equal(t, types.Money(10), model.Expense)
				assert.Equal(t, types.Money(10), model.MonthlyWithoutCredit)

			},
		},
		{
			name: "add value in montlhy with credit",
			dto: messaging.MessagingDto{
				Amount:       10,
				Movement:     messaging.TypeExpense,
				MovementType: string(helpers.ExpenseTypeMensal),
				WithCredit:   true,
			},
			model:      models.Aggregated{},
			multiplier: 1,
			check: func(t *testing.T, model models.Aggregated) {
				assert.Equal(t, types.Money(10), model.Expense)
				assert.Equal(t, types.Money(10), model.MonthlyWithCredit)

			},
		},
		{
			name: "add value in invoice",
			dto: messaging.MessagingDto{
				Amount:       10,
				Movement:     messaging.TypeExpense,
				MovementType: string(helpers.ExpenseTypeFatura),
			},
			model:      models.Aggregated{},
			multiplier: 1,
			check: func(t *testing.T, model models.Aggregated) {
				assert.Equal(t, types.Money(10), model.Expense)
				assert.Equal(t, types.Money(10), model.Invoice)

			},
		},
		{
			name: "add value in investment",
			dto: messaging.MessagingDto{
				Amount:   10,
				Movement: messaging.TypeInvestment,
			},
			model:      models.Aggregated{},
			multiplier: 1,
			check: func(t *testing.T, model models.Aggregated) {
				assert.Equal(t, types.Money(10), model.Investments)
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			model := tt.model
			calculateValuesModel(&tt.dto, &model, tt.multiplier)
			tt.check(t, model)
		})
	}

}
