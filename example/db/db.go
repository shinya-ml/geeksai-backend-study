package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Open() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=example sslmode=disable", "db", 5432, "root", "p@ssword"))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
