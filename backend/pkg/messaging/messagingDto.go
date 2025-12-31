package messaging

import (
	"errors"
	"time"

	"github.com/GeuberLucas/Gofre/backend/pkg/types"
)

type Movement string
type ActionType string

const (
	TypeIncome     Movement = "INCOME"
	TypeExpense    Movement = "EXPENSE"
	TypeInvestment Movement = "INVESTMENT"

	ActionInsert ActionType = "CREATE"
	ActionUpdate ActionType = "UPDATE"
	ActionDelete ActionType = "DELETE"
)

func (m Movement) IsValid() bool {
	switch m {
	case TypeIncome, TypeExpense, TypeInvestment:
		return true
	}
	return false
}
func (act ActionType) IsValid() bool {
	switch act {
	case ActionInsert, ActionUpdate, ActionDelete:
		return true
	}
	return false
}

type MessagingDto struct {
	Month               time.Month  `json:"month"`
	MonthOld            time.Month  `json:"month_old"`
	YearOld             uint        `json:"year_old"`
	Year                uint        `json:"year"`
	Amount              types.Money `json:"amount"`
	AmountOld           types.Money `json:"old_amount"`
	Movement            Movement    `json:"movement"`
	MovementTypeOld     string      `json:"movement_type_old"`
	MovementType        string      `json:"movement_type"`
	MovementCategoryOld string      `json:"movement_category_old"`
	MovementCategory    string      `json:"movement_category"`
	WithCredit          bool        `json:"with_credit"`
	WithCreditOld       bool        `json:"with_credit_old"`
	IsConfirmedOld      bool        `json:"is_confirmed_old"`
	IsConfirmed         bool        `json:"is_confirmed"`
	Action              ActionType  `json:"action"`
	UserId              int         `json:"user_id"`
}

func (md *MessagingDto) IsValid() error {
	if md.Month == 0 || md.Month > 12 {
		return errors.New("Messaging: validate struct: Invalid Month")
	}
	if md.Year < 1 {
		return errors.New("Messaging: validate struct: Invalid Year")
	}
	if !md.Movement.IsValid() {
		return errors.New("Messaging: validate struct: Invalid Movement")
	}
	if !md.Action.IsValid() {
		return errors.New("Messaging: validate struct: Invalid Action")
	}
	if len(md.MovementType) == 0 {
		return errors.New("Messaging: validate struct: Invalid Movement")
	}
	if md.Movement == TypeExpense && len(md.MovementCategory) == 0 {
		return errors.New("Messaging: validate struct: MovementCategory is required for Expense")
	}
	if md.Action == ActionUpdate && md.AmountOld == 0 {
		return errors.New("Messaging: validate struct: Amount Old is required for update")
	}
	return nil
}
