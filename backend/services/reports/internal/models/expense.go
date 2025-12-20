package models

import "github.com/GeuberLucas/Gofre/backend/pkg/types"

type Expense struct {
	Month    int
	Year     int
	Planned  types.Money
	Actual   types.Money
	Pending  types.Money
	Invoice  types.Money
	Variable types.Money
	Monthly  types.Money
	UserId   int
}
