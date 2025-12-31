package repository

import (
	"database/sql"

	"github.com/GeuberLucas/Gofre/backend/pkg/helpers"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/interfaces"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/models"
)

type RevenueRepository struct {
	db *sql.DB
}

func NewRevenueRepository(conn *sql.DB) interfaces.IReportsRepository[models.Revenue] {
	return &RevenueRepository{db: conn}
}

func (inv *RevenueRepository) InsertOrUpdate(tx *sql.Tx, model *models.Revenue) (helpers.ErrorType, error) {
	sqlCommand := `INSERT INTO reports.revenue (month, year, user_id, planned, actual, pending) 
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
func (inv *RevenueRepository) GetAll(userId int) ([]models.Revenue, helpers.ErrorType, error) {
	var data []models.Revenue
	var sqlCommand string = `SELECT month, year, user_id, planned, actual, pending
    FROM reports.revenue
    WHERE user_id = $1
    `

	rows, err := inv.db.Query(sqlCommand)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}
	defer rows.Close()

	for rows.Next() {
		var Revenue models.Revenue
		err := rows.Scan(&Revenue.Month, &Revenue.Year, &Revenue.UserId, &Revenue.Planned, &Revenue.Actual, &Revenue.Pending)
		if err != nil {
			return nil, helpers.INTERNAL, err
		}
		data = append(data, Revenue)
	}
	return data, helpers.NONE, nil
}
func (inv *RevenueRepository) GetByMonthAndYear(userId int, month int, year int) (models.Revenue, helpers.ErrorType, error) {
	var Revenue models.Revenue
	var sqlCommand string = `
    SELECT month, year, user_id, planned, actual, pending
    FROM reports.revenue
    WHERE user_id = $1 AND month = $2 AND year = $3`

	row := inv.db.QueryRow(sqlCommand, month, year)
	err := row.Scan(&Revenue.Month, &Revenue.Year, &Revenue.UserId, &Revenue.Planned, &Revenue.Actual, &Revenue.Pending)
	if err != nil {
		return Revenue, helpers.INTERNAL, err
	}
	return Revenue, helpers.NONE, nil
}
func (inv *RevenueRepository) GetByYear(userId int, year int) ([]models.Revenue, helpers.ErrorType, error) {
	var data []models.Revenue
	var sqlCommand string = `
    SELECT month, year, user_id, planned, actual, pending
    FROM reports.revenue
    WHERE user_id = $1 AND year = $2
    `

	rows, err := inv.db.Query(sqlCommand, year)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}
	defer rows.Close()

	for rows.Next() {
		var Revenue models.Revenue
		err := rows.Scan(
			&Revenue.Month, &Revenue.Year, &Revenue.UserId, &Revenue.Planned, &Revenue.Actual, &Revenue.Pending,
		)
		if err != nil {
			return nil, helpers.INTERNAL, err
		}
		data = append(data, Revenue)
	}
	return data, helpers.NONE, nil
}
