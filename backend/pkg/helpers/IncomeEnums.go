package helpers

// IncomeType define a origem da receita
type IncomeType string

const (
	IncomeTypeTrabalho      IncomeType = "Trabalho"
	IncomeTypeExtra         IncomeType = "Extra"
	IncomeTypeInvestimento  IncomeType = "Investimento"
	IncomeTypeAposentadoria IncomeType = "Aposentadoria"
	IncomeTypeResgate       IncomeType = "Resgate"
	IncomeTypeOutros        IncomeType = "Outros"
)
