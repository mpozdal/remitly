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

type AddBankPayload struct {
	SwiftCode     string `json:"swiftCode"`
	BankName      string `json:"bankName"`
	Address       string `json:"address"`
	CountryISO2   string `json:"countryISO2"`
	IsHeadquarter bool   `json:"isHeadquarter"`
	CountryName   string `json:"countryName"`
}

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type ReponseByCountry struct {
	CountryISO2 string         `json:"countryISO2"`
	CountryName string         `json:"countryName"`
	Data        []BankResponse `json:"swiftCodes"`
}

type BankResponse struct {
	SwiftCode     string `json:"swiftCode"`
	BankName      string `json:"bankName"`
	Address       string `json:"address"`
	CountryISO2   string `json:"countryISO2"`
	CountryName   string `json:"countryName"`
	IsHeadquarter bool   `json:"isHeadquarter"`
}

type BankWithBranchesReponse struct {
	SwiftCode     string         `json:"swiftCode"`
	BankName      string         `json:"bankName"`
	Address       string         `json:"address"`
	CountryISO2   string         `json:"countryISO2"`
	CountryName   string         `json:"countryName"`
	IsHeadquarter bool           `json:"isHeadquarter"`
	Branches      []BankResponse `json:"branches,omitempty"`
}