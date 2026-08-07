package expense

type IExpenseRepository interface {
	Create(model Expense) error
	GetAll(userId uint) ([]Expense, error)
	GetById(id uint) (Expense, error)
	Update(model Expense) error
	Delete(id uint, userId uint) error
}
