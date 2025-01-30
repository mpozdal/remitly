package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type DBManager struct {
	DB *sql.DB
}

func NewDBManager(cfg mysql.Config) (*DBManager, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return &DBManager{DB: db}, nil
}
