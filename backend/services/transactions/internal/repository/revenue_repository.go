package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GeuberLucas/Gofre/backend/pkg/helpers"
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/models"
)

type IRevenueRepository interface {
	Create(model models.Revenue) error
	GetAll() ([]models.Revenue, error)
	GetById(id int64) (models.Revenue, error)
	GetByUserId(userId int64) ([]models.Revenue, error)
	Update(model models.Revenue) error
	Delete(id int64, userId int64) error
}

type RevenueRepository struct {
	db *sql.DB
}

func NewRevenueRepository(db *sql.DB) *RevenueRepository {
	return &RevenueRepository{db: db}
}

func (r RevenueRepository) Create(model models.Revenue) error {
	sqlCommand := `INSERT INTO transactions.revenue(
	user_id, description, origin, type, received_date, is_recieved,amount)
	VALUES ($1, $2, $3, $4, $5, $6,$7);`

	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(
		model.UserId,
		model.Description,
		model.Origin,
		model.Type.ToDBString(),
		model.ReceiveDate,
		model.IsRecieved,
		model.Amount,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r RevenueRepository) GetAll() ([]models.Revenue, error) {
	var sqlCommand string = `SELECT id, user_id, description, origin, type, received_date, is_recieved,amount
	FROM transactions.revenue;`

	rows, err := r.db.Query(sqlCommand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var revenues []models.Revenue
	for rows.Next() {
		var revenue models.Revenue
		var typeStr string
		err := rows.Scan(
			&revenue.ID,
			&revenue.UserId,
			&revenue.Description,
			&revenue.Origin,
			&typeStr,
			&revenue.ReceiveDate,
			&revenue.IsRecieved,
			&revenue.Amount,
		)
		if err != nil {
			return nil, err
		}
		revenue.Type = helpers.ParseIncomeType(typeStr)
		revenues = append(revenues, revenue)
	}
	return revenues, nil

}
func (r RevenueRepository) GetById(id int64) (models.Revenue, error) {
	var revenue models.Revenue
	var typeStr string
	var sqlCommand string = `SELECT id, user_id, description, origin, type, received_date, is_recieved,amount
	FROM transactions.revenue
	WHERE id=$1;`

	row := r.db.QueryRow(sqlCommand, id)
	err := row.Scan(
		&revenue.ID,
		&revenue.UserId,
		&revenue.Description,
		&revenue.Origin,
		&typeStr,
		&revenue.ReceiveDate,
		&revenue.IsRecieved,
		&revenue.Amount,
	)
	if err != nil {
		return revenue, err
	}
	revenue.Type = helpers.ParseIncomeType(typeStr)
	return revenue, nil
}
func (r RevenueRepository) GetByUserId(userId int64) ([]models.Revenue, error) {

	var sqlCommand string = `SELECT id, user_id, description, origin, type, received_date, is_recieved,amount
	FROM transactions.revenue
	WHERE user_id=$1;`

	rows, err := r.db.Query(sqlCommand, userId)
	if err != nil {
		return nil, fmt.Errorf("repository:revenue:sql query: %s ", err)
	}
	defer rows.Close()

	var revenues []models.Revenue
	for rows.Next() {
		var revenue models.Revenue
		var typeStr string

		err := rows.Scan(
			&revenue.ID,
			&revenue.UserId,
			&revenue.Description,
			&revenue.Origin,
			&typeStr,
			&revenue.ReceiveDate,
			&revenue.IsRecieved,
			&revenue.Amount,
		)
		if err != nil {
			return nil, fmt.Errorf("repository:revenue:scan rows: %s", err)
		}

		// Conversão string -> Enum
		revenue.Type = helpers.ParseIncomeType(typeStr)
		revenues = append(revenues, revenue)
	}
	return revenues, nil
}
func (r RevenueRepository) Update(model models.Revenue) error {
	var sqlCommand string = `UPDATE transactions.revenue
	SET
	description=$1, 
	origin=$2,
	type=$3,
	received_date=$4,
	is_recieved=$5,
	amount=$6
	WHERE id=$7;`

	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(
		model.Description,
		model.Origin,
		model.Type.ToDBString(),
		model.ReceiveDate,
		model.IsRecieved,
		model.Amount,
		model.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
func (r RevenueRepository) Delete(id int64, userId int64) error {
	var sqlCommand string = "DELETE FROM transactions.revenue where id=$1 and user_id=$2;"

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
