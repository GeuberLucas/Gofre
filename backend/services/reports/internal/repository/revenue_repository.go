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

func (inv *RevenueRepository) InsertOrUpdate(model *models.Revenue) (helpers.ErrorType, error) {
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
func (inv *RevenueRepository) GetAll(userId int) ([]models.Revenue, helpers.ErrorType, error) {
	var data []models.Revenue
	var sqlCommand string = ``

	rows, err := inv.db.Query(sqlCommand)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}
	defer rows.Close()

	for rows.Next() {
		var Revenue models.Revenue
		err := rows.Scan(&Revenue.Month, &Revenue.Year, &Revenue.Actual, &Revenue.Pending, &Revenue.Planned, &Revenue.UserId)
		if err != nil {
			return nil, helpers.INTERNAL, err
		}
		data = append(data, Revenue)
	}
	return data, helpers.NONE, nil
}
func (inv *RevenueRepository) GetByMonthAndYear(userId int, month int, year int) (models.Revenue, helpers.ErrorType, error) {
	var Revenue models.Revenue
	var sqlCommand string = ``

	row := inv.db.QueryRow(sqlCommand, month, year)
	err := row.Scan(&Revenue.Month, &Revenue.Year, &Revenue.Actual, &Revenue.Pending, &Revenue.Planned, &Revenue.UserId)
	if err != nil {
		return Revenue, helpers.INTERNAL, err
	}
	return Revenue, helpers.NONE, nil
}
func (inv *RevenueRepository) GetByYear(userId int, year int) ([]models.Revenue, helpers.ErrorType, error) {
	var data []models.Revenue
	var sqlCommand string = ``

	rows, err := inv.db.Query(sqlCommand, year)
	if err != nil {
		return nil, helpers.INTERNAL, err
	}
	defer rows.Close()

	for rows.Next() {
		var Revenue models.Revenue
		err := rows.Scan(&Revenue.Month, &Revenue.Year, &Revenue.Actual, &Revenue.Pending, &Revenue.Planned, &Revenue.UserId)
		if err != nil {
			return nil, helpers.INTERNAL, err
		}
		data = append(data, Revenue)
	}
	return data, helpers.NONE, nil
}
