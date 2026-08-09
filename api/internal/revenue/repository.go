package revenue

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
)

type IRevenueRepository interface {
	Create(model Revenue) error
	GetAll(userId uint) ([]Revenue, error)
	GetById(id uint) (Revenue, error)
	Update(model Revenue) error
	Delete(id uint, userId uint) error
}

type RevenueRepository struct {
	db *sql.DB
}

func NewRevenueRepository(db *sql.DB) IRevenueRepository {
	return &RevenueRepository{db: db}
}

func (r RevenueRepository) Create(model Revenue) error {
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

func (r RevenueRepository) GetById(id uint) (Revenue, error) {
	var revenue Revenue
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
func (r RevenueRepository) GetAll(userId uint) ([]Revenue, error) {

	var sqlCommand string = `SELECT id, user_id, description, origin, type, received_date, is_recieved,amount
	FROM transactions.revenue
	WHERE user_id=$1;`

	rows, err := r.db.Query(sqlCommand, userId)
	if err != nil {
		return nil, fmt.Errorf("repository:revenue:sql query: %s ", err)
	}
	defer rows.Close()

	var revenues []Revenue
	for rows.Next() {
		var revenue Revenue
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
func (r RevenueRepository) Update(model Revenue) error {
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
func (r RevenueRepository) Delete(id uint, userId uint) error {
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
