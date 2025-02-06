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

func NewDBManagerWithCon(db *sql.DB) (*DBManager, error) {
	return &DBManager{DB: db}, nil
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

func (db *DBManager) DeleteBank(swiftCode string) (bool, error) {

	res, err := db.DB.Exec(DELETE_BANK_QUERY, swiftCode)
	if err != nil {
		return false, fmt.Errorf("error deleting bank: %s", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("error: %s", err)
	}
	return rowsAffected > 0, nil
}

func (db *DBManager) GetCountry(countryISO2 string) (*types.Country, error) {

	rows, err := db.DB.Query(GET_COUNTRY_QUERY, countryISO2)
	if err != nil {
		return nil, fmt.Errorf("error getting country by countryISO2: %s", err)
	}
	country := new(types.Country)
	for rows.Next() {
		err := rows.Scan(&country.CountryISO2, &country.CountryName)
		if err != nil {
			return nil, fmt.Errorf("error getting country by countryISO2: %s", err)
		}
	}
	return country, nil
}

func (db *DBManager) GetBanksByCountry(countryISO2code string) ([]types.Bank, error) {

	var banks []types.Bank
	rows, err := db.DB.Query(GET_BANKS_BY_COUNTRY_QUERY, countryISO2code)
	if err != nil {
		return nil, fmt.Errorf("error getting banks by countryISO2: %s", err)
	}
	found := false
	for rows.Next() {
		var bank types.Bank
		err := rows.Scan(&bank.SwiftCode, &bank.BankName, &bank.Address, &bank.CountryISO2, &bank.CountryName, &bank.IsHeadquarter, &sql.NullString{Valid: false})
		if err != nil {
			return nil, fmt.Errorf("error scanning rows: %s", err)
		}
		banks = append(banks, bank)
		found = true
	}
	if !found {
		return nil, fmt.Errorf("no banks found for country: %s", countryISO2code)
	}
	return banks, nil

}

func (db *DBManager) GetBankBySwiftCode(swiftCode string) (*types.Bank, error) {
	bank := new(types.Bank)

	err := db.DB.QueryRow(GET_BANK_BY_SWIFTCODE_QUERY, swiftCode).Scan(&bank.SwiftCode, &bank.BankName, &bank.Address, &bank.CountryISO2, &bank.CountryName, &bank.IsHeadquarter, &bank.HeadquarterSwiftCode)
	if err != nil {
		return nil, fmt.Errorf("error getting bank by swift code: %s", err)
	}
	return bank, nil
}
func (db *DBManager) GetBranchesByHQSwiftCode(swiftCode string) ([]types.Bank, error) {
	var branches []types.Bank

	rows, err := db.DB.Query(GET_BRANCHES_BY_HQ_SWIFTCODE_QUERY, swiftCode)
	if err != nil {
		return nil, fmt.Errorf("error getting bank by hq swift code: %s", err)
	}
	defer rows.Close()
	for rows.Next() {
		var branch types.Bank
		err := rows.Scan(&branch.SwiftCode, &branch.BankName, &branch.Address, &branch.CountryISO2, &branch.CountryName, &branch.IsHeadquarter, &branch.HeadquarterSwiftCode)
		if err != nil {
			return nil, fmt.Errorf("error getting branches by swift code: %s", err)
		}
		branches = append(branches, branch)
	}
	return branches, nil
}
