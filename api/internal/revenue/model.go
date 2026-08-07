package revenue

import (
	"errors"
	"time"

	dtos "github.com/GeuberLucas/Gofre/api/pkg/DTOs"
	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
	"github.com/GeuberLucas/Gofre/api/pkg/types"
)

type Revenue struct {
	ID          uint
	UserId      uint
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
func ToModel(re dtos.RevenueDto) Revenue {
	amount := types.FloatToMoney(re.Amount)
	return Revenue{
		ID:          uint(re.ID),
		UserId:      uint(re.UserId),
		Description: re.Description,
		Origin:      re.Origin,
		Type:        re.Type,
		ReceiveDate: re.ReceiveDate,
		IsRecieved:  re.IsRecieved,
		Amount:      amount,
	}
}
