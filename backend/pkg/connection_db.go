package pkg

import (
	"database/sql"

	_ "github.com/lib/pq"
) 

func ConnectToDatabase() (*sql.DB, error) {
	stringConnection := "user=postgres password=123456 host=localhost port=5432 dbname=Gofre sslmode=disable"
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