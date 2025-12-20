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

func (inv *InvestmentsRepository) InsertOrUpdate(model *models.Investment) (helpers.ErrorType, error) {
	sqlCommand := ``

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
func (inv *InvestmentsRepository) GetAll(userId int) ([]models.Investment, helpers.ErrorType, error) {
	var data []models.Investment
	var sqlCommand string = ``

	rows, err := inv.db.Query(sqlCommand)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}
	defer rows.Close()

	for rows.Next() {
		var Investment models.Investment
		err := rows.Scan(&Investment.Month, &Investment.Year, &Investment.Actual, &Investment.Pending, &Investment.Planned, &Investment.UserId)
		if err != nil {
			return nil, helpers.INTERNAL, err
		}
		data = append(data, Investment)
	}
	return data, helpers.NONE, nil
}
func (inv *InvestmentsRepository) GetByMonthAndYear(userId int, month int, year int) (models.Investment, helpers.ErrorType, error) {
	var Investment models.Investment
	var sqlCommand string = ``

	row := inv.db.QueryRow(sqlCommand, month, year)
	err := row.Scan(&Investment.Month, &Investment.Year, &Investment.Actual, &Investment.Pending, &Investment.Planned, &Investment.UserId)
	if err != nil {
		return Investment, helpers.INTERNAL, err
	}
	return Investment, helpers.NONE, nil
}
func (inv *InvestmentsRepository) GetByYear(userId int, year int) ([]models.Investment, helpers.ErrorType, error) {
	var data []models.Investment
	var sqlCommand string = ``

	rows, err := inv.db.Query(sqlCommand, year)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}
	defer rows.Close()

	for rows.Next() {
		var Investment models.Investment
		err := rows.Scan(&Investment.Month, &Investment.Year, &Investment.Actual, &Investment.Pending, &Investment.Planned, &Investment.UserId)
		if err != nil {
			return nil, helpers.INTERNAL, err
		}
		data = append(data, Investment)
	}
	return data, helpers.NONE, nil
}
