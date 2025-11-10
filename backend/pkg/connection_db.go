package pkg

import (
	"database/sql"

	_ "github.com/lib/pq"
) 

func ConnectToDatabase() (*sql.DB, error) {
	stringConnection := "postgres:123456@tcp(localhost:5432)/Gofre"
	db, error := sql.Open("postgres", stringConnection)
	if error != nil {
		return nil, error
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}


func CloseDatabaseConnection(db *sql.DB) error {
	return db.Close()
}