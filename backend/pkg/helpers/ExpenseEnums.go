package helpers

// ExpenseCategory representa as categorias de despesas
type ExpenseCategory string

const (
	ExpenseCategoryMercadoGeral     ExpenseCategory = "mercado geral"
	ExpenseCategoryDelivery         ExpenseCategory = "delivery"
	ExpenseCategoryRestauranteBares ExpenseCategory = "restaurante e bares"
	ExpenseCategoryVestuario        ExpenseCategory = "vestuário"
	ExpenseCategoryMoradia          ExpenseCategory = "moradia"
	ExpenseCategoryUtilidades       ExpenseCategory = "utilidades"
	ExpenseCategoryDecoracao        ExpenseCategory = "decoração"
	ExpenseCategoryEducacao         ExpenseCategory = "educação"
	ExpenseCategoryDependentes      ExpenseCategory = "dependentes"
	ExpenseCategorySaude            ExpenseCategory = "saúde"
	ExpenseCategoryEntretenimento   ExpenseCategory = "entretenimento"
	ExpenseCategoryServicos         ExpenseCategory = "serviços"
	ExpenseCategoryImpostos         ExpenseCategory = "impostos"
	ExpenseCategoryTransporte       ExpenseCategory = "transporte"
	ExpenseCategoryPresentes        ExpenseCategory = "presentes"
	ExpenseCategoryPets             ExpenseCategory = "pets"
	ExpenseCategoryViagens          ExpenseCategory = "viagens"
	ExpenseCategoryDoacoes          ExpenseCategory = "doações"
	ExpenseCategoryApostas          ExpenseCategory = "apostas"
	ExpenseCategoryLivre            ExpenseCategory = "livre"
	ExpenseCategoryOutros           ExpenseCategory = "outros"
)

// ExpenseType define a recorrência ou tipo da despesa
type ExpenseType string

const (
	ExpenseTypeMensal   ExpenseType = "mensal"
	ExpenseTypeVariavel ExpenseType = "variável"
	ExpenseTypeFatura   ExpenseType = "fatura"
)

// PaymentMethod define os meios de pagamento
type PaymentMethod string

const (
	PaymentMethodPix      PaymentMethod = "pix"
	PaymentMethodDebito   PaymentMethod = "debito"
	PaymentMethodCredito  PaymentMethod = "credito"
	PaymentMethodBoleto   PaymentMethod = "boleto"
	PaymentMethodDinheiro PaymentMethod = "dinheiro"
	PaymentMethodTED      PaymentMethod = "ted"
	PaymentMethodCheque   PaymentMethod = "cheque"
)
