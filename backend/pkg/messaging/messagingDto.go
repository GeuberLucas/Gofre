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
	Month            time.Month  `json:"month"`
	Year             uint        `json:"year"`
	Amount           types.Money `json:"amount"`
	AmountOld        types.Money `json:"old_amount"`
	Movement         Movement    `json:"movement"`
	MovementType     string      `json:"movement_type"`
	MovementCategory string      `json:"movement_category"`
	WithCredit       bool        `json:"with_credit"`
	IsConfirmed      bool        `json:"is_confirmed"`
	Action           ActionType  `json:"action"`
	UserId           int         `json:"user_id"`
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

func NewMessagingDto(month time.Month,
	year uint,
	amount types.Money,
	amountOld types.Money,
	movement Movement,
	movementType string,
	movementCategory string,
	withCredit bool,
	isConfirmed bool,
	action ActionType,
	userId int) (MessagingDto, error) {
	messagingDto := MessagingDto{
		Month:            month,
		Year:             year,
		Amount:           amount,
		Movement:         movement,
		MovementType:     movementType,
		MovementCategory: movementCategory,
		WithCredit:       withCredit,
		IsConfirmed:      isConfirmed,
		Action:           action,
		UserId:           userId,
		AmountOld:        amountOld,
	}
	if err := messagingDto.IsValid(); err != nil {
		return MessagingDto{}, err
	}
	return messagingDto, nil

}
