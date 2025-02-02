package types

import "database/sql"

type Bank struct {
	SwiftCode            string         `json:"swiftCode"`
	BankName             string         `json:"bankName"`
	Address              string         `json:"address"`
	CountryISO2          string         `json:"countryISO2"`
	CountryName          string         `json:"countryName"`
	IsHeadquarter        bool           `json:"isHeadquarter"`
	HeadquarterSwiftCode sql.NullString `json:"headquarterSwiftCode,omitempty"`
}
type Country struct {
	CountryISO2 string `json:"countryISO2"`
	CountryName string `json:"countryName"`
}