package helpers

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

// 1. ExpenseCategory
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

// ToDBString converte o inteiro da categoria para a string esperada pelo banco
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
