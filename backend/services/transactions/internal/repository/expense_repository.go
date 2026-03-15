package repository

import (
	"database/sql"
	"log"

	"github.com/GeuberLucas/Gofre/backend/pkg/helpers"
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/models"
)

type IExpenseRepository interface {
	Create(model models.Expense) error
	GetAll() ([]models.Expense, error)
	GetById(id int64) (models.Expense, error)
	GetByUserId(userId int64) ([]models.Expense, error)
	Update(model models.Expense) error
	Delete(id int64, userId int64) error
}

type ExpenseRepository struct {
	db *sql.DB
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{db: db}
}

func (r ExpenseRepository) Create(model models.Expense) error {
	sqlCommand := `INSERT INTO transactions.expenses(
	user_id, description, target, category,amount, type, payment_method, payment_date, is_paid)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8,$9);`

	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(
		model.UserId,
		model.Description,
		model.Target,
		model.Category.ToDBString(),
		model.Amount,
		model.Type.ToDBString(),
		model.PaymentMethod.ToDBString(),
		model.PaymentDate,
		model.IsPaid,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r ExpenseRepository) GetAll() ([]models.Expense, error) {
	var sqlCommand string = `SELECT id, user_id, description, target, category,amount, type, payment_method, payment_date, is_paid
	FROM transactions.expenses;`

	rows, err := r.db.Query(sqlCommand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var expense models.Expense
		var categoryStr, typeStr, paymentMethodStr string
		err := rows.Scan(
			&expense.ID,
			&expense.UserId,
			&expense.Description,
			&expense.Target,
			&categoryStr,
			&expense.Amount,
			&typeStr,
			&paymentMethodStr,
			&expense.PaymentDate,
			&expense.IsPaid,
		)
		if err != nil {
			return nil, err
		}
		expense.Category = helpers.ParseExpenseCategory(categoryStr)
		expense.Type = helpers.ParseExpenseType(typeStr)
		expense.PaymentMethod = helpers.ParsePaymentMethod(paymentMethodStr)
		expenses = append(expenses, expense)
	}
	return expenses, nil
}
func (r ExpenseRepository) GetById(id int64) (models.Expense, error) {
	var expense models.Expense
	var categoryStr, typeStr, paymentMethodStr string
	var sqlCommand string = `SELECT id, user_id, description, target, category,amount, type, payment_method, payment_date, is_paid
	FROM transactions.expenses
	WHERE id=$1;`

	row := r.db.QueryRow(sqlCommand, id)
	err := row.Scan(
		&expense.ID,
		&expense.UserId,
		&expense.Description,
		&expense.Target,
		&categoryStr,
		&expense.Amount,
		&typeStr,
		&paymentMethodStr,
		&expense.PaymentDate,
		&expense.IsPaid,
	)
	if err != nil {
		return expense, err
	}
	expense.Category = helpers.ParseExpenseCategory(categoryStr)
	expense.Type = helpers.ParseExpenseType(typeStr)
	expense.PaymentMethod = helpers.ParsePaymentMethod(paymentMethodStr)
	return expense, nil
}
func (r ExpenseRepository) GetByUserId(userId int64) ([]models.Expense, error) {
	var sqlCommand string = `SELECT id, user_id, description, target, category,amount, type, payment_method, payment_date, is_paid
	FROM transactions.expenses
	WHERE user_id=$1;`

	rows, err := r.db.Query(sqlCommand, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var expense models.Expense
		var categoryStr, typeStr, paymentMethodStr string
		err := rows.Scan(
			&expense.ID,
			&expense.UserId,
			&expense.Description,
			&expense.Target,
			&categoryStr,
			&expense.Amount,
			&typeStr,
			&paymentMethodStr,
			&expense.PaymentDate,
			&expense.IsPaid,
		)
		if err != nil {
			return nil, err
		}
		expense.Category = helpers.ParseExpenseCategory(categoryStr)
		expense.Type = helpers.ParseExpenseType(typeStr)
		expense.PaymentMethod = helpers.ParsePaymentMethod(paymentMethodStr)
		expenses = append(expenses, expense)
	}
	return expenses, nil
}
func (r ExpenseRepository) Update(model models.Expense) error {
	var sqlCommand string = `UPDATE transactions.expenses
	SET description=$1, target=$2, category=$3,amount=$4, type=$5, payment_method=$6, payment_date=$7, is_paid=$8
	WHERE id=$9;`

	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(
		model.Description,
		model.Target,
		model.Category.ToDBString(),
		model.Amount,
		model.Type.ToDBString(),
		model.PaymentMethod.ToDBString(),
		model.PaymentDate,
		model.IsPaid,
		model.ID,
	)
	if err != nil {
		return err
	}
	return nil

}
func (r ExpenseRepository) Delete(id int64, userId int64) error {
	var sqlCommand string = "DELETE FROM transactions.expenses where id=$1 and user_id=$2;"

	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(id, userId)
	if err != nil {
		return err
	}
	log.Println("User deleted successfully")
	return nil
}
