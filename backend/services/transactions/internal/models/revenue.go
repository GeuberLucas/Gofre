package models

import (
	"errors"
	"time"

	"github.com/GeuberLucas/Gofre/backend/pkg/types"
)

type Revenue struct {
	ID          int64
	UserId      int64
	Description string
	Origin      string
	Type        string
	ReceiveDate time.Time
	Amount      types.Money
	IsRecieved  bool
}

func (re Revenue) Isvalid() error {

	if re.UserId == 0 {
		return errors.New("revenue:validate:UserId required")
	}
	if re.Origin == "" {
		return errors.New("revenue:validate:Origin required")
	}

	if re.Type == "" {
		return errors.New("revenue:validate:Type required")
	}
	if re.ReceiveDate.IsZero() {
		return errors.New("revenue:validate:ReceiveDate required")
	}

	if re.Amount <= 0 {
		return errors.New("revenue:validate:Amount not be equal or minor than zero")
	}

	return nil
}
