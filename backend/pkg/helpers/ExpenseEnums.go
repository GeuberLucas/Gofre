package helpers

// ExpenseCategory representa as categorias de despesas
type ExpenseCategory string

const (
	ExpenseCategoryMercadoGeral     ExpenseCategory = "Mercado geral"
	ExpenseCategoryDelivery         ExpenseCategory = "Delivery"
	ExpenseCategoryRestauranteBares ExpenseCategory = "Restaurante e bares"
	ExpenseCategoryVestuario        ExpenseCategory = "Vestuário"
	ExpenseCategoryMoradia          ExpenseCategory = "Moradia"
	ExpenseCategoryUtilidades       ExpenseCategory = "Utilidades"
	ExpenseCategoryDecoracao        ExpenseCategory = "Decoração"
	ExpenseCategoryEducacao         ExpenseCategory = "Educação"
	ExpenseCategoryDependentes      ExpenseCategory = "Dependentes"
	ExpenseCategorySaude            ExpenseCategory = "Saúde"
	ExpenseCategoryEntretenimento   ExpenseCategory = "Entretenimento"
	ExpenseCategoryServicos         ExpenseCategory = "Serviços"
	ExpenseCategoryImpostos         ExpenseCategory = "Impostos"
	ExpenseCategoryTransporte       ExpenseCategory = "Transporte"
	ExpenseCategoryPresentes        ExpenseCategory = "Presentes"
	ExpenseCategoryPets             ExpenseCategory = "Pets"
	ExpenseCategoryViagens          ExpenseCategory = "Viagens"
	ExpenseCategoryDoacoes          ExpenseCategory = "Doações"
	ExpenseCategoryApostas          ExpenseCategory = "Apostas"
	ExpenseCategoryLivre            ExpenseCategory = "Livre"
	ExpenseCategoryOutros           ExpenseCategory = "Outros"
)

// ExpenseType define a recorrência ou tipo da despesa
type ExpenseType string

const (
	ExpenseTypeMensal   ExpenseType = "Mensal"
	ExpenseTypeVariavel ExpenseType = "Variável"
	ExpenseTypeFatura   ExpenseType = "Fatura"
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
