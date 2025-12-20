package models

import "github.com/GeuberLucas/Gofre/backend/pkg/types"

type Revenue struct {
	Month   int
	Year    int
	Planned types.Money
	Actual  types.Money
	Pending types.Money
	UserId  int
}
