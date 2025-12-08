package models

import "slices"

// asset represents a type of investment
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
