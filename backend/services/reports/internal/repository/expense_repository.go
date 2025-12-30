package repository

import (
	"database/sql"

	"github.com/GeuberLucas/Gofre/backend/pkg/helpers"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/interfaces"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/models"
)

type ExpensesRepository struct {
	db *sql.DB
}

func NewExpensesRepository(conn *sql.DB) interfaces.IReportsRepository[models.Expense] {
	return &ExpensesRepository{db: conn}
}

func (inv *ExpensesRepository) InsertOrUpdate(model *models.Expense) (helpers.ErrorType, error) {
	sqlCommand := `INSERT INTO reports.expense (month, year, user_id, planned, actual, pending, invoice, variable, monthly) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
ON CONFLICT (month, year, user_id) 
DO UPDATE SET 
    planned = EXCLUDED.planned,
    actual = EXCLUDED.actual,
    pending = EXCLUDED.pending,
    invoice = EXCLUDED.invoice,
    variable = EXCLUDED.variable,
    monthly = EXCLUDED.monthly;`

	statement, err := inv.db.Prepare(sqlCommand)
	if err != nil {
		return helpers.INTERNAL, err
	}
	defer statement.Close()
	_, err = statement.Exec(model.Month, model.Year, model.Actual, model.Pending, model.Planned, model.UserId)

	if err != nil {
		return helpers.INTERNAL, err
	}

	return helpers.NONE, nil
}
func (inv *ExpensesRepository) GetAll(userId int) ([]models.Expense, helpers.ErrorType, error) {
	var data []models.Expense
	var sqlCommand string = ``

	rows, err := inv.db.Query(sqlCommand)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}
	defer rows.Close()

	for rows.Next() {
		var Expense models.Expense
		err := rows.Scan(&Expense.Month, &Expense.Year, &Expense.Actual, &Expense.Pending, &Expense.Planned, &Expense.UserId)
		if err != nil {
			return nil, helpers.INTERNAL, err
		}
		data = append(data, Expense)
	}
	return data, helpers.NONE, nil
}
func (inv *ExpensesRepository) GetByMonthAndYear(userId int, month int, year int) (models.Expense, helpers.ErrorType, error) {
	var Expense models.Expense
	var sqlCommand string = ``

	row := inv.db.QueryRow(sqlCommand, month, year)
	err := row.Scan(&Expense.Month, &Expense.Year, &Expense.Actual, &Expense.Pending, &Expense.Planned, &Expense.UserId)
	if err != nil {
		return Expense, helpers.INTERNAL, err
	}
	return Expense, helpers.NONE, nil
}
func (inv *ExpensesRepository) GetByYear(userId int, year int) ([]models.Expense, helpers.ErrorType, error) {
	var data []models.Expense
	var sqlCommand string = ``

	rows, err := inv.db.Query(sqlCommand, year)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}
	defer rows.Close()

	for rows.Next() {
		var Expense models.Expense
		err := rows.Scan(&Expense.Month, &Expense.Year, &Expense.Actual, &Expense.Pending, &Expense.Planned, &Expense.UserId)
		if err != nil {
			return nil, helpers.INTERNAL, err
		}
		data = append(data, Expense)
	}
	return data, helpers.NONE, nil
}
