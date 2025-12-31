package repository

import (
	"database/sql"

	"github.com/GeuberLucas/Gofre/backend/pkg/helpers"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/interfaces"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/models"
)

type InvestmentsRepository struct {
	db *sql.DB
}

func NewInvestmentsRepository(conn *sql.DB) interfaces.IReportsRepository[models.Investment] {
	return &InvestmentsRepository{db: conn}
}

func (inv *InvestmentsRepository) InsertOrUpdate(tx *sql.Tx, model *models.Investment) (helpers.ErrorType, error) {
	sqlCommand := `INSERT INTO reports.investments (month, year, user_id, planned, actual, pending) 
    VALUES ($1, $2, $3, $4, $5, $6) 
    ON CONFLICT (month, year, user_id) 
    DO UPDATE SET 
        planned = EXCLUDED.planned,
        actual = EXCLUDED.actual,
        pending = EXCLUDED.pending;`

	statement, err := inv.db.Prepare(sqlCommand)
	if err != nil {
		return helpers.INTERNAL, err
	}
	defer statement.Close()
	_, err = statement.Exec(
		model.Month,   // $1
		model.Year,    // $2
		model.UserId,  // $3
		model.Planned, // $4
		model.Actual,  // $5
		model.Pending, // $6
	)

	if err != nil {
		return helpers.INTERNAL, err
	}

	return helpers.NONE, nil
}
func (inv *InvestmentsRepository) GetAll(userId int) ([]models.Investment, helpers.ErrorType, error) {
	var data []models.Investment
	var sqlCommand string = `SELECT month, year, user_id, planned, actual, pending
    FROM reports.investments
    WHERE user_id = $1
    `

	rows, err := inv.db.Query(sqlCommand)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Investment
		err := rows.Scan(&i.Month, &i.Year, &i.UserId, &i.Planned, &i.Actual, &i.Pending)
		if err != nil {
			return nil, helpers.INTERNAL, err
		}
		data = append(data, i)
	}
	return data, helpers.NONE, nil
}
func (inv *InvestmentsRepository) GetByMonthAndYear(userId int, month int, year int) (models.Investment, helpers.ErrorType, error) {
	var Investment models.Investment
	var sqlCommand string = `
    SELECT month, year, user_id, planned, actual, pending
    FROM reports.investments
    WHERE user_id = $1 AND month = $2 AND year = $3`

	row := inv.db.QueryRow(sqlCommand, userId, month, year)
	err := row.Scan(
		&Investment.Month, &Investment.Year, &Investment.UserId, &Investment.Planned, &Investment.Actual, &Investment.Pending,
	)
	if err != nil {
		return Investment, helpers.INTERNAL, err
	}
	return Investment, helpers.NONE, nil
}
func (inv *InvestmentsRepository) GetByYear(userId int, year int) ([]models.Investment, helpers.ErrorType, error) {
	var data []models.Investment
	var sqlCommand string = `
    SELECT month, year, user_id, planned, actual, pending
    FROM reports.investments
    WHERE user_id = $1 AND year = $2
    `

	rows, err := inv.db.Query(sqlCommand, userId, year)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Investment
		err := rows.Scan(
			&i.Month, &i.Year, &i.UserId, &i.Planned, &i.Actual, &i.Pending,
		)
		if err != nil {
			return nil, helpers.INTERNAL, err
		}
		data = append(data, i)
	}
	return data, helpers.NONE, nil
}
