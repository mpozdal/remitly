package db

import (
	"database/sql"
	"fmt"
	"log"
	"mpozdal/remitly/types"

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

func (db *DBManager) AddCountry(country types.Country) error {
	_, err := db.DB.Exec(CREATE_COUNTRY_QUERY, country.CountryISO2, country.CountryName)
	if err != nil {
		return fmt.Errorf("error inserting new country: %s", err)
	}
	return nil

}
func (db *DBManager) AddBank(bank types.Bank) (bool, error) {

	res, err := db.DB.Exec(CREATE_BANK_QUERY, bank.SwiftCode, bank.BankName, bank.Address, bank.CountryISO2, bank.IsHeadquarter, bank.HeadquarterSwiftCode)
	if err != nil {
		return false, fmt.Errorf("error inserting new bank: %s", err)

	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("error: %s", err)
	}
	return rowsAffected > 0, nil
}
