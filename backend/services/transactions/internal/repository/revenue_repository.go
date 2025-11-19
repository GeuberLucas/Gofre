package repository

import (
	"database/sql"
	"log"

	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/models"
)

type IRevenueRepository interface{
	Create( model models.Revenue) error
	GetAll() ([]models.Revenue,error)
	GetById(id int)(models.Revenue,error)
	GetByUserId(userId int)(models.Revenue,error)
	Update(model models.Revenue) error
	Delete(id int) error
}

type RevenueRepository struct {
	db *sql.DB
}

func NewRevenueRepository(db *sql.DB) *RevenueRepository{
	return  &RevenueRepository{db:db}
}

func(r RevenueRepository) Create( model models.Revenue) error{
	sqlCommand := `INSERT INTO transactions.revenue(
	user_id, description, origin, type, received_date, is_recieved)
	VALUES ($1, $2, $3, $4, $5, $6);`
	

	statement, err := r.db.Prepare(sqlCommand)
	defer statement.Close()
	if err != nil {
		return err
	}
	_,err = statement.Exec(model.UserId,model.Description,model.Origin,model.Type,model.ReceiveDate,model.IsRecieved)
	if err != nil {
		return err
	}	
	return nil
}



func(r RevenueRepository) GetAll() ([]models.Revenue,error){
	var sqlCommand string=`SELECT id, user_id, description, origin, type, received_date, is_recieved
	FROM transactions.revenue;`
	

	rows, err := r.db.Query(sqlCommand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var revenues []models.Revenue
	for rows.Next() {
		var revenue models.Revenue
		err := rows.Scan(&revenue.ID,&revenue.UserId,&revenue.Description,&revenue.Origin,&revenue.Type,&revenue.ReceiveDate,&revenue.IsRecieved )
		if err != nil {
			return nil, err
		}
		revenues = append(revenues, revenue)
	}
	return revenues, nil

}
func(r RevenueRepository) GetById(id int)(models.Revenue,error){
	var revenue models.Revenue
	var sqlCommand string=`SELECT id, user_id, description, origin, type, received_date, is_recieved
	FROM transactions.revenue
	WHERE id=$1;`
	
	row := r.db.QueryRow(sqlCommand, id)
	err := row.Scan(&revenue.ID,&revenue.UserId,&revenue.Description,&revenue.Origin,&revenue.Type,&revenue.ReceiveDate,&revenue.IsRecieved)
	if err != nil {
		return revenue, err
	}
	return revenue, nil
}
func(r RevenueRepository) GetByUserId(userId int)(models.Revenue,error){
	var revenue models.Revenue
	var sqlCommand string=`SELECT id, user_id, description, origin, type, received_date, is_recieved
	FROM transactions.revenue
	WHERE user_id=$1;`
	
	row := r.db.QueryRow(sqlCommand, userId)
	err := row.Scan(&revenue.ID,&revenue.UserId,&revenue.Description,&revenue.Origin,&revenue.Type,&revenue.ReceiveDate,&revenue.IsRecieved)
	if err != nil {
		return revenue, err
	}
	return revenue, nil
}
func(r RevenueRepository) Update(model models.Revenue) error{
	var sqlCommand string=`UPDATE transactions.revenue
	SET  user_id=$1,
	description=$2, 
	origin=$3,
	type=$4,
	received_date=$5,
	is_recieved=$6
	WHERE id=$7;`
	
	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(model.UserId,model.Description,model.Origin,model.Type,model.ReceiveDate,model.IsRecieved)
	if err != nil {
		return err
	}
	log.Println("User updated successfully")
	return nil 
}
func(r RevenueRepository) Delete(id int) error{
	var sqlCommand string="DELETE FROM transactions.revenue where id=$1;"
	
	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(id)
	if err != nil {		log.Fatal(err)
		return err
	}
	log.Println("User deleted successfully")
	return nil
}