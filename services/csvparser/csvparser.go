package csvparser

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"

	"mpozdal/remitly/types"
)

type CSVPareser struct {
	FilePath string
}

func NewCSVParser(filePath string) *CSVPareser {
	return &CSVPareser{FilePath: filePath}
}

func (cp *CSVPareser) ReadRecords() ([][]string, error) {
	log.Println(cp.FilePath)
	file, err := os.Open(cp.FilePath)
	if err != nil {
		log.Fatal("Error opening file")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading file")
	}

	return records, nil

}

func getHeadquarters(records [][]string) map[string]string {
	headquartersMap := make(map[string]string)
	for _, record := range records[1:] {
		swiftCode := record[1]
		if len(swiftCode) >= 8 && swiftCode[8:] == "XXX" {
			headquartersMap[swiftCode[:8]] = swiftCode
		}
	}
	return headquartersMap
}

func (cp *CSVPareser) ParseRecords() ([]types.Country, []types.Bank, error) {
	records, err := cp.ReadRecords()
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	headquartersMap := getHeadquarters(records)
	countriesMap := make(map[string]types.Country)
	var banks []types.Bank

	for _, record := range records[1:] {
		countryISO2 := record[0]
		countryName := record[6]
		countriesMap[countryISO2] = types.Country{
			CountryISO2: countryISO2,
			CountryName: countryName,
		}

		swiftCode := record[1]
		isHeadquarter := len(swiftCode) >= 8 && swiftCode[8:] == "XXX"
		var headquarterSwiftCode = sql.NullString{
			Valid: false,
		}

		if !isHeadquarter {
			if len(swiftCode) >= 8 {
				prefix := swiftCode[:8]
				if hqSwiftCode, exists := headquartersMap[prefix]; exists {
					headquarterSwiftCode =
						sql.NullString{
							String: hqSwiftCode,
							Valid:  true,
						}
				}

			}
		}
		bank := types.Bank{
			SwiftCode:            swiftCode,
			BankName:             record[3],
			Address:              record[4],
			CountryISO2:          countryISO2,
			IsHeadquarter:        isHeadquarter,
			HeadquarterSwiftCode: headquarterSwiftCode,
		}
		banks = append(banks, bank)
	}

	var countries []types.Country
	for _, country := range countriesMap {
		countries = append(countries, country)
	}

	return countries, banks, nil

}
