package investments

import (
	"errors"
	"slices"
	"time"

	"github.com/GeuberLucas/Gofre/api/pkg/types"
)

type Asset struct {
	id   uint
	name string
}

var assets = []Asset{

	{1, "Títulos privados"},
	{2, "Títulos públicos"},
	{3, "Ações"},
	{4, "ETFs"},
	{5, "FIIs"},
	{6, "Fundos"},
	{7, "Commodities"},
	{8, "Derivativos"},
	{9, "Criptomoeda"},
	{10, "Exterior"},
	{11, "Poupança"},
	{12, "Outros"},
}

func GetAssetName(Id uint) string {
	idx := slices.IndexFunc(assets, func(a Asset) bool { return a.id == Id })
	return assets[idx].name
}

type Portfolio struct {
	Id           uint
	User_id      uint
	Asset_id     uint
	Deposit_date time.Time
	Broker       string
	Amount       types.Money
	IsDone       bool
	Description  string
}

func (p *Portfolio) IsValid() error {
	if p.Deposit_date.IsZero() {
		return errors.New("Portfolio struct: Validate: deposit date is required")
	}
	if p.Asset_id == 0 {
		return errors.New("Portfolio struct: Validate: asset type is required")
	}

	return nil
}
