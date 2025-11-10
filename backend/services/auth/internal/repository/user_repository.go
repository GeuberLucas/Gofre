package repository

import (
	"log"
	"time"

	"github.com/GeuberLucas/Gofre/backend/pkg"
	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/models"
)

func NewUserRepository( user models.User)(int64) {
	var sqlCommand string="insert into auth.users (name,lastName,cell_phone,username, email,password,created_at,updated_at) values (?,?,?,?,?,?,?,?)"
	dbConn,error := pkg.ConnectToDatabase()
	if error != nil {
		log.Fatal(error)
		return 0
	}
	defer pkg.CloseDatabaseConnection(dbConn)

	statement, err := dbConn.Prepare(sqlCommand)
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
func GetUsers() ([]models.User, error) {
	var sqlCommand string="select id, username, name, lastName, cell_phone, email, password, created_at, updated_at from auth.users"
	dbConn,error := pkg.ConnectToDatabase()
	if error != nil {
		log.Fatal(error)
		return nil, error
	}
	defer pkg.CloseDatabaseConnection(dbConn)

	rows, err := dbConn.Query(sqlCommand)
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

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	var sqlCommand string="select id, username, name, lastName, cell_phone, email, password, created_at, updated_at from auth.users where username = ?"
	dbConn,error := pkg.ConnectToDatabase()
	if error != nil {
		log.Fatal(error)
		return user, error
	}
	defer pkg.CloseDatabaseConnection(dbConn)
	row := dbConn.QueryRow(sqlCommand, username)
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.LastName, &user.Cellphone, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByID(id int64) (models.User, error) {
	var user models.User
	var sqlCommand string="select id, username, name, lastName, cell_phone, email, password, created_at, updated_at from auth.users where id = ?"
	dbConn,error := pkg.ConnectToDatabase()	
	if error != nil {
		log.Fatal(error)
		return user, error
	}
	defer pkg.CloseDatabaseConnection(dbConn)
	row := dbConn.QueryRow(sqlCommand, id)
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.LastName, &user.Cellphone, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(user models.User)  {
	var sqlCommand string="update auth.users set name=?, lastName=?, cell_phone=?, username=?, email=?, password=?, updated_at=? where id=?"
	dbConn,error := pkg.ConnectToDatabase()
	if error != nil {
		log.Fatal(error)
		return 
	}
	defer pkg.CloseDatabaseConnection(dbConn)
	statement, err := dbConn.Prepare(sqlCommand)
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

func DeleteUser(id int64)  {
	var sqlCommand string="delete from auth.users where id=?"
	dbConn,error := pkg.ConnectToDatabase()
	if error != nil {
		log.Fatal(error)
		return
	}
	defer pkg.CloseDatabaseConnection(dbConn)
	statement, err := dbConn.Prepare(sqlCommand)
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
