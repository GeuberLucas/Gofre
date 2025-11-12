package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/models"
)

type UserRepository struct {
	db *sql.DB
}


func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser( user models.User)(int64) {
	sqlCommand := `INSERT INTO auth.users (name, last_name, cell_phone, username, email, password, created_at, updated_at)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id
					`
	

	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	
	var idLastInsert int64
	err = statement.QueryRow(user.Name, user.LastName, user.Cellphone, user.Username, user.Email, user.Password, time.Now(), time.Now()).Scan(&idLastInsert)
	if err != nil {
		log.Fatal(err)
		return 0
	}	
	defer statement.Close()
	return idLastInsert
}
func (r *UserRepository) GetUsers() ([]models.User, error) {
	var sqlCommand string="select id, username, name, last_name, cell_phone, email, password, created_at, updated_at from auth.users"
	

	rows, err := r.db.Query(sqlCommand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.LastName, &user.Cellphone, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	var sqlCommand string="select id, username, name, last_name, cell_phone, email, password, created_at, updated_at from auth.users where username = $1"
	
	row := r.db.QueryRow(sqlCommand, username)
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.LastName, &user.Cellphone, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	var sqlCommand string="select id, username, name, last_name, cell_phone, email, password, created_at, updated_at from auth.users where email = $1"
	
	row := r.db.QueryRow(sqlCommand, email)
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.LastName, &user.Cellphone, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByID(id int64) (models.User, error) {
	var user models.User
	var sqlCommand string="select id, username, name, last_name, cell_phone, email, password, created_at, updated_at from auth.users where id = $1"
	
	row := r.db.QueryRow(sqlCommand, id)
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.LastName, &user.Cellphone, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user models.User)  {
	var sqlCommand string="update auth.users set name=$1, last_name=$2, cell_phone=$3, username=$4, email=$5, password=$6, updated_at=$7 where id=$8"
	
	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		log.Fatal(err)
		return 
	}
	defer statement.Close()
	_, err = statement.Exec(user.Name, user.LastName, user.Cellphone, user.Username, user.Email, user.Password, time.Now(), user.ID)
	if err != nil {
		log.Fatal(err)
		return 
	}
	log.Println("User updated successfully")
	
}

func (r *UserRepository) UpdateUserPassword(userId int64,password []byte){
	var sqlCommand string="update auth.users set  password=$1 ,updated_at=$2 where id=$3"
	
	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		log.Fatal(err)
		return 
	}
	defer statement.Close()
	_, err = statement.Exec(password, time.Now(), userId)
	if err != nil {
		log.Fatal(err)
		return 
	}
	log.Println("User updated successfully")
}

func (r *UserRepository) DeleteUser(id int64)  {
	var sqlCommand string="delete from auth.users where id=$1"
	
	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer statement.Close()
	_, err = statement.Exec(id)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("User deleted successfully")
}
