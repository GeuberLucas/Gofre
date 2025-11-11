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
	var sqlCommand string="insert into auth.users (name,lastName,cell_phone,username, email,password,created_at,updated_at) values (?,?,?,?,?,?,?,?)"
	

	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	defer statement.Close()
	
	lastInsert, err := statement.Exec(user.Name, user.LastName, user.Cellphone, user.Username, user.Email, user.Password, time.Now(), time.Now())
	if err != nil {
		log.Fatal(err)
		return 0
	}
	log.Println("User inserted successfully")
	
	idLastInsert, err := lastInsert.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return 0
	}

	return idLastInsert
}
func (r *UserRepository) GetUsers() ([]models.User, error) {
	var sqlCommand string="select id, username, name, lastName, cell_phone, email, password, created_at, updated_at from auth.users"
	

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
	var sqlCommand string="select id, username, name, lastName, cell_phone, email, password, created_at, updated_at from auth.users where username = ?"
	
	row := r.db.QueryRow(sqlCommand, username)
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.LastName, &user.Cellphone, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByID(id int64) (models.User, error) {
	var user models.User
	var sqlCommand string="select id, username, name, lastName, cell_phone, email, password, created_at, updated_at from auth.users where id = ?"
	
	row := r.db.QueryRow(sqlCommand, id)
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.LastName, &user.Cellphone, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user models.User)  {
	var sqlCommand string="update auth.users set name=?, lastName=?, cell_phone=?, username=?, email=?, password=?, updated_at=? where id=?"
	
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

func (r *UserRepository) DeleteUser(id int64)  {
	var sqlCommand string="delete from auth.users where id=?"
	
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
