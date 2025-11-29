package models_test

import (
	"errors"
	"testing"
	"time"

	"github.com/GeuberLucas/Gofre/backend/pkg/types"
	. "github.com/GeuberLucas/Gofre/backend/services/transaction/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestValidationRevenue(t *testing.T) {
	base := Revenue{
		ID:          1,
		UserId:      1,
		Description: "valid description",
		Origin:      "valid",
		Type:        "valid",
		ReceiveDate: time.Date(2025, 10, 25, 0, 0, 0, 0, time.UTC),
		Amount:      types.FloatToMoney(50),
		IsRecieved:  true,
	}
	type testCase struct {
		modify  func(e *Revenue)
		wantErr error
	}

	t.Run("test validation revenue model", func(t *testing.T) {
		tests := []testCase{
			{modify: func(e *Revenue) {}, wantErr: nil},
			{modify: func(e *Revenue) { e.UserId = 0 }, wantErr: errors.New("revenue:validate:UserId required")},
			{modify: func(e *Revenue) { e.Origin = "" }, wantErr: errors.New("revenue:validate:Origin required")},
			{modify: func(e *Revenue) { e.Type = "" }, wantErr: errors.New("revenue:validate:Type required")},
			{modify: func(e *Revenue) { e.ReceiveDate = time.Time{} }, wantErr: errors.New("revenue:validate:ReceiveDate required")},
			{modify: func(e *Revenue) { e.Amount = types.Money(-1) }, wantErr: errors.New("revenue:validate:Amount not be equal or minor than zero")},
			{modify: func(e *Revenue) { e.Amount = types.Money(0) }, wantErr: errors.New("revenue:validate:Amount not be equal or minor than zero")},
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
