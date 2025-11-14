package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
) 

func ConnectToDatabase() (*sql.DB, error) {

	stringConnection := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
	os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_HOST"),os.Getenv("DB_PORT"),os.Getenv("DB_DBNAME"),
)
	db, err:= sql.Open("postgres", stringConnection)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}


func CloseDatabaseConnection(db *sql.DB) error {
	return db.Close()
}