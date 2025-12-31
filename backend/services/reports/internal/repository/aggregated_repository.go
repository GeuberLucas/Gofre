package repository

import (
	"database/sql"

	"github.com/GeuberLucas/Gofre/backend/pkg/helpers"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/interfaces"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/models"
)

type AggregatedRepository struct {
	db *sql.DB
}

func NewAggregatedRepository(conn *sql.DB) interfaces.IReportsRepository[models.Aggregated] {
	return &AggregatedRepository{db: conn}
}

func (agr *AggregatedRepository) InsertOrUpdate(model *models.Aggregated) (helpers.ErrorType, error) {
	sqlCommand := `INSERT INTO reports.aggregated (
	  "month", 
	  "year", 
	  "user_id", 
	  "revenue", 
	  "expense", 
	  "investments", 
	  "monthly_without_credit", 
	  "monthly_with_credit", 
	  "variable_without_credit", 
	  "variable_with_credit", 
	  "invoice", 
	  "result"
	) 
	VALUES (
	  $1,   -- month
	  $2,   -- year
	  $3,   -- user_id
	  $4,   -- revenue
	  $5,   -- expense
	  $6,   -- investments
	  $7,   -- monthly_without_credit
	  $8,   -- monthly_with_credit
	  $9,   -- variable_without_credit
	  $10,  -- variable_with_credit
	  $11,  -- invoice
	  $12   -- result
	)
	ON CONFLICT ("month", "year", "user_id") 
	DO UPDATE SET
	  "revenue"                 = EXCLUDED."revenue",
	  "expense"                 = EXCLUDED."expense",
	  "investments"             = EXCLUDED."investments",
	  "monthly_without_credit"  = EXCLUDED."monthly_without_credit",
	  "monthly_with_credit"     = EXCLUDED."monthly_with_credit",
	  "variable_without_credit" = EXCLUDED."variable_without_credit",
	  "variable_with_credit"    = EXCLUDED."variable_with_credit",
	  "invoice"                 = EXCLUDED."invoice",
	  "result"                  = EXCLUDED."result";`

	statement, err := agr.db.Prepare(sqlCommand)
	if err != nil {
		return helpers.INTERNAL, err
	}
	defer statement.Close()
	_, err = statement.Exec(
		model.Month,
		model.Year,
		model.UserId,
		model.Revenue,
		model.Expense,
		model.Investments,
		model.MonthlyWithoutCredit,
		model.MonthlyWithCredit,
		model.VariableWithoutCredit,
		model.VariableWithCredit,
		model.Invoice,
		model.Result,
	)

	if err != nil {
		return helpers.INTERNAL, err
	}

	return helpers.NONE, nil
}
func (agr *AggregatedRepository) GetAll(userId int) ([]models.Aggregated, helpers.ErrorType, error) {
	var data []models.Aggregated
	var sqlCommand string = `SELECT month, 
		year, 
		revenue, 
		expense, 
		investments, 
		monthly_without_credit,
		monthly_with_credit,
		variable_without_credit,
		variable_with_credit,
		invoice,
		result,
		user_id
	FROM reports.aggregated where user_id=$1;`

	rows, err := agr.db.Query(sqlCommand, userId)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}
	defer rows.Close()

	for rows.Next() {
		var aggregated models.Aggregated
		err := rows.Scan(&aggregated.Month, &aggregated.Year, &aggregated.Revenue, &aggregated.Expense, &aggregated.Investments,
			&aggregated.MonthlyWithoutCredit, &aggregated.MonthlyWithCredit, &aggregated.VariableWithoutCredit,
			&aggregated.VariableWithoutCredit, &aggregated.Invoice, &aggregated.Result, &aggregated.UserId)
		if err != nil {
			return nil, helpers.INTERNAL, err
		}
		data = append(data, aggregated)
	}
	return data, helpers.NONE, nil
}
func (agr *AggregatedRepository) GetByMonthAndYear(userId int, month int, year int) (models.Aggregated, helpers.ErrorType, error) {
	var aggregated models.Aggregated
	var sqlCommand string = `SELECT month, 
		year, 
		revenue, 
		expense, 
		investments, 
		monthly_without_credit,
		monthly_with_credit,
		variable_without_credit,
		variable_with_credit,
		invoice,
		result,
		user_id
	FROM reports.aggregated where user_id=$1 and month= $2 and year=$3;`

	row := agr.db.QueryRow(sqlCommand, userId, month, year)
	err := row.Scan(&aggregated.Month, &aggregated.Year, &aggregated.Revenue, &aggregated.Expense, &aggregated.Investments,
		&aggregated.MonthlyWithoutCredit, &aggregated.MonthlyWithCredit, &aggregated.VariableWithoutCredit,
		&aggregated.VariableWithoutCredit, &aggregated.Invoice, &aggregated.Result, &aggregated.UserId)
	if err != nil {
		return aggregated, helpers.INTERNAL, err
	}
	return aggregated, helpers.NONE, nil
}
func (agr *AggregatedRepository) GetByYear(userId int, year int) ([]models.Aggregated, helpers.ErrorType, error) {
	var data []models.Aggregated
	var sqlCommand string = `SELECT month, 
		year, 
		revenue, 
		expense, 
		investments, 
		monthly_without_credit,
		monthly_with_credit,
		variable_without_credit,
		variable_with_credit,
		invoice,
		result,
		user_id
	FROM reports.aggregated where user_id=$1 and year=$2;`

	rows, err := agr.db.Query(sqlCommand, userId, year)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}
	defer rows.Close()

	for rows.Next() {
		var aggregated models.Aggregated
		err := rows.Scan(&aggregated.Month, &aggregated.Year, &aggregated.Revenue, &aggregated.Expense, &aggregated.Investments,
			&aggregated.MonthlyWithoutCredit, &aggregated.MonthlyWithCredit, &aggregated.VariableWithoutCredit,
			&aggregated.VariableWithoutCredit, &aggregated.Invoice, &aggregated.Result, &aggregated.UserId)
		if err != nil {
			return nil, helpers.INTERNAL, err
		}
		data = append(data, aggregated)
	}
	return data, helpers.NONE, nil
}
