package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Connection driver
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "123456"
	dbname   = "aulas-go"
)

// Connect open connection with database
func Connect() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
