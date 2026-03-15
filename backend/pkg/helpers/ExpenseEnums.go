package helpers

import "strings"

// ExpenseType define a recorrência ou tipo da despesa
type ExpenseType int

const (
	ExpenseTypeMensal ExpenseType = iota
	ExpenseTypeVariavel
	ExpenseTypeFatura
)

func (t ExpenseType) ToDBString() string {
	dbValues := []string{
		"Mensal",
		"Variável",
		"Fatura",
	}
	if t < 0 || int(t) >= len(dbValues) {
		return "Outros"
	}
	return dbValues[t]
}
func ParseExpenseType(s string) ExpenseType {
	switch s {
	case "Mensal":
		return ExpenseTypeMensal
	case "Variável":
		return ExpenseTypeVariavel
	case "Fatura":
		return ExpenseTypeFatura
	default:
		return -1
	}
}

type ExpenseCategory int

const (
	ExpenseCategoryMercadoGeral ExpenseCategory = iota // 0
	ExpenseCategoryDelivery
	ExpenseCategoryRestauranteBares
	ExpenseCategoryVestuario
	ExpenseCategoryMoradia
	ExpenseCategoryUtilidades
	ExpenseCategoryDecoracao
	ExpenseCategoryEducacao
	ExpenseCategoryDependentes
	ExpenseCategorySaude
	ExpenseCategoryEntretenimento
	ExpenseCategoryServicos
	ExpenseCategoryImpostos
	ExpenseCategoryTransporte
	ExpenseCategoryPresentes
	ExpenseCategoryPets
	ExpenseCategoryViagens
	ExpenseCategoryDoacoes
	ExpenseCategoryApostas
	ExpenseCategoryLivre
	ExpenseCategoryOutros
)

func (e ExpenseCategory) ToDBString() string {
	dbValues := []string{
		"Mercado geral",
		"Delivery",
		"Restaurante e bares",
		"Vestuário",
		"Moradia",
		"Utilidades",
		"Decoração",
		"Educação",
		"Dependentes",
		"Saúde",
		"Entretenimento",
		"Serviços",
		"Impostos",
		"Transporte",
		"Presentes",
		"Pets",
		"Viagens",
		"Doações",
		"Apostas",
		"Livre",
		"Outros",
	}

	if e < 0 || int(e) >= len(dbValues) {
		return "Outros"
	}

	return dbValues[e]
}
func ParseExpenseCategory(s string) ExpenseCategory {
	switch s {
	case "Mercado geral":
		return ExpenseCategoryMercadoGeral
	case "Delivery":
		return ExpenseCategoryDelivery
	case "Restaurante e bares":
		return ExpenseCategoryRestauranteBares
	case "Vestuário":
		return ExpenseCategoryVestuario
	case "Moradia":
		return ExpenseCategoryMoradia
	case "Utilidades":
		return ExpenseCategoryUtilidades
	case "Decoração":
		return ExpenseCategoryDecoracao
	case "Educação":
		return ExpenseCategoryEducacao
	case "Dependentes":
		return ExpenseCategoryDependentes
	case "Saúde":
		return ExpenseCategorySaude
	case "Entretenimento":
		return ExpenseCategoryEntretenimento
	case "Serviços":
		return ExpenseCategoryServicos
	case "Impostos":
		return ExpenseCategoryImpostos
	case "Transporte":
		return ExpenseCategoryTransporte
	case "Presentes":
		return ExpenseCategoryPresentes
	case "Pets":
		return ExpenseCategoryPets
	case "Viagens":
		return ExpenseCategoryViagens
	case "Doações":
		return ExpenseCategoryDoacoes
	case "Apostas":
		return ExpenseCategoryApostas
	case "Livre":
		return ExpenseCategoryLivre
	case "Outros":
		return ExpenseCategoryOutros
	default:
		return ExpenseCategoryOutros
	}
}

// 2. PaymentMethod
type PaymentMethod int

const (
	PaymentMethodPix PaymentMethod = iota
	PaymentMethodDebito
	PaymentMethodCredito
	PaymentMethodBoleto
	PaymentMethodDinheiro
	PaymentMethodTED
	PaymentMethodCheque
)

// ToDBString converte o inteiro do método de pagamento para a string esperada pelo banco
func (p PaymentMethod) ToDBString() string {
	dbValues := []string{
		"pix",
		"debito",
		"credito",
		"boleto",
		"dinheiro",
		"ted",
		"cheque",
	}

	if p < 0 || int(p) >= len(dbValues) {
		return ""
	}

	return dbValues[p]
}
func ParsePaymentMethod(s string) PaymentMethod {
	// Usamos strings.ToLower para garantir que "PIX", "Pix" ou "pix" funcionem
	switch strings.ToLower(s) {
	case "pix":
		return PaymentMethodPix
	case "debito", "débito":
		return PaymentMethodDebito
	case "credito", "crédito":
		return PaymentMethodCredito
	case "boleto":
		return PaymentMethodBoleto
	case "dinheiro":
		return PaymentMethodDinheiro
	case "ted":
		return PaymentMethodTED
	case "cheque":
		return PaymentMethodCheque
	default:
		return -1 // Valor inválido
	}
}
