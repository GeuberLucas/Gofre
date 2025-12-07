package messaging

import (
	"errors"

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
	Month            uint        `json:"month"`
	Year             uint        `json:"year"`
	Amount           types.Money `json:"amount"`
	Movement         Movement    `json:"movement"`
	MovementType     string      `json:"movement_type"`
	MovementCategory string      `json:"movement_category"`
	WithCredit       bool        `json:"with_credit"`
	IsConfirmed      bool        `json:"is_confirmed"`
	Action           ActionType  `json:"action"`
}

func (md *MessagingDto) IsValid() error {
	if md.Month == 0 || md.Month > 12 {
		return errors.New("Messaging: validade struct: Invalid Month")
	}
	if md.Year < 1 {
		return errors.New("Messaging: validade struct: Invalid Year")
	}
	if !md.Movement.IsValid() {
		return errors.New("Messaging: validade struct: Invalid Movement")
	}
	if !md.Action.IsValid() {
		return errors.New("Messaging: validade struct: Invalid Action")
	}
	if len(md.MovementType) == 0 {
		return errors.New("Messaging: validade struct: Invalid Movement")
	}
	if md.Movement == TypeExpense && len(md.MovementCategory) == 0 {
		return errors.New("Messaging: validade struct: MovementCategory is required for Expense")
	}
	if md.Movement == TypeExpense && len(md.MovementCategory) == 0 {
		return errors.New("Messaging: validade struct: MovementCategory is required for Expense")
	}
	return nil
}

func NewMessagingDto(month uint,
	year uint,
	amount types.Money,
	movement Movement,
	movementType string,
	movementCategory string,
	withCredit bool,
	isConfirmed bool,
	action ActionType) (*MessagingDto, error) {
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
	}
	if err := messagingDto.IsValid(); err != nil {
		return nil, err
	}
	return &messagingDto, nil

}
