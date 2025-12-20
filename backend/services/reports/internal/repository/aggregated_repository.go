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
	sqlCommand := ``

	statement, err := agr.db.Prepare(sqlCommand)
	if err != nil {
		return helpers.INTERNAL, err
	}
	defer statement.Close()
	_, err = statement.Exec(model.Month, model.Year, model.Revenue, model.Expense, model.Investments, model.MonthlyWithoutCredit, model.MonthlyWithCredit, model.VariableWithoutCredit, model.VariableWithoutCredit, model.Invoice, model.Result, model.UserId)

	if err != nil {
		return helpers.INTERNAL, err
	}

	return helpers.NONE, nil
}
func (agr *AggregatedRepository) GetAll(userId int) ([]models.Aggregated, helpers.ErrorType, error) {
	var data []models.Aggregated
	var sqlCommand string = ``

	rows, err := agr.db.Query(sqlCommand)
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
	var sqlCommand string = ``

	row := agr.db.QueryRow(sqlCommand, month, year)
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
	var sqlCommand string = ``

	rows, err := agr.db.Query(sqlCommand, year)
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
