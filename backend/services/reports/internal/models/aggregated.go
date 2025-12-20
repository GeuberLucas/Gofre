package models

import "github.com/GeuberLucas/Gofre/backend/pkg/types"

type Aggregated struct {
	Month                 int
	Year                  int
	Revenue               types.Money
	Expense               types.Money
	Investments           types.Money
	MonthlyWithCredit     types.Money
	MonthlyWithoutCredit  types.Money
	VariableWithoutCredit types.Money
	VariableWithCredit    types.Money
	Invoice               types.Money
	Result                types.Money
	UserId                int
}
