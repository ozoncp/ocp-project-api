package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	dbDriver  = "postgres"
	dbUser    = "lobanov"
	dbName    = "ocp"
	dbSslmode = "disable"
)

func OpenDB() (*sqlx.DB, error) {
	db, err := sqlx.Open(dbDriver, fmt.Sprintf("user=%s dbname=%s sslmode=%s", dbUser, dbName, dbSslmode))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect(dbDriver, fmt.Sprintf("user=%s dbname=%s sslmode=%s", dbUser, dbName, dbSslmode))
	if err != nil {
		return nil, err
	}

	return db, nil
}
