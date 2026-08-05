package helpers

// IncomeType define a origem da receita
type IncomeType int

const (
	IncomeTypeTrabalho IncomeType = iota
	IncomeTypeExtra
	IncomeTypeInvestimento
	IncomeTypeAposentadoria
	IncomeTypeResgate
	IncomeTypeOutros
)

func (i IncomeType) ToDBString() string {
	dbValues := []string{
		"Trabalho",
		"Extra",
		"Investimento",
		"Aposentadoria",
		"Resgate",
		"Outros",
	}

	// Se o valor for inválido, retornamos "Outros" por segurança
	if i < 0 || int(i) >= len(dbValues) {
		return "Outros"
	}

	return dbValues[i]
}

func ParseIncomeType(s string) IncomeType {
	switch s {
	case "Trabalho":
		return IncomeTypeTrabalho
	case "Extra":
		return IncomeTypeExtra
	case "Investimento":
		return IncomeTypeInvestimento
	case "Aposentadoria":
		return IncomeTypeAposentadoria
	case "Resgate":
		return IncomeTypeResgate
	case "Outros":
		return IncomeTypeOutros
	default:
		return IncomeTypeOutros
	}
}
