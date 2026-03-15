package models

import (
	"errors"
	"time"

	"github.com/GeuberLucas/Gofre/backend/pkg/helpers"
	"github.com/GeuberLucas/Gofre/backend/pkg/types"
)

type Revenue struct {
	ID          int64
	UserId      int64
	Description string
	Origin      string
	Type        helpers.IncomeType
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

	if re.Type < 0 || re.Type > helpers.IncomeTypeOutros {
		return errors.New("revenue:validate:Type invalid")
	}
	if re.ReceiveDate.IsZero() {
		return errors.New("revenue:validate:ReceiveDate required")
	}

	if re.Amount <= 0 {
		return errors.New("revenue:validate:Amount not be equal or minor than zero")
	}

	return nil
}
