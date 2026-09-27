package profile

import "github.com/GeuberLucas/Gofre/api/pkg/types"

type Profile struct {
	ID             uint
	UserId         uint
	CompleteName   string
	CellPhone      string
	InitialBalance types.Money
}
