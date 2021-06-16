package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-project-api/internal/config"
)

func OpenDB() (*sqlx.DB, error) {
	db, err := sqlx.Open(
		config.Global.DB.Driver,
		fmt.Sprintf(
			"user=%s dbname=%s sslmode=%s",
			config.Global.DB.User,
			config.Global.DB.Name,
			config.Global.DB.Sslmode),
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect(
		config.Global.DB.Driver,
		fmt.Sprintf(
			"user=%s dbname=%s sslmode=%s",
			config.Global.DB.User,
			config.Global.DB.Name,
			config.Global.DB.Sslmode),
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
