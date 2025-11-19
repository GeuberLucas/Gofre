package repository

import (
	"database/sql"

	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/models"
)

type IExpenseRepository interface {
	Create(model models.Expense) error
	GetAll() ([]models.Expense, error)
	GetById(id int) (models.Expense, error)
	GetByUserId(userId int) (models.Expense, error)
	Update(model models.Expense) error
	Delete(id int) error
}

type ExpenseRepository struct {
	db *sql.DB
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{db: db}
}

func (r ExpenseRepository) Create(model models.Expense) error {
	sqlCommand := `INSERT INTO transactions.expenses(
	user_id, description, target, category, type, payment_method, payment_date, is_paid)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`

	statement, err := r.db.Prepare(sqlCommand)
	defer statement.Close()
	if err != nil {
		return err
	}
	
	_,err = statement.Exec(model.UserId,model.Description,model.Target,model.Category,model.Type,model.PaymentMethod,model.PaymentDate,model.IsPaid)
	if err != nil {
		return err
	}	
	
	return nil
}

func (r ExpenseRepository) GetAll(userId int) ([]models.Expense, error) {}
func (r ExpenseRepository) GetById(id int) (models.Expense, error)      {}
func (r ExpenseRepository) Update(model models.Expense) error           {}
func (r ExpenseRepository) Delete(id int) error                         {}