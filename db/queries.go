package db

const (
	CREATE_COUNTRY_QUERY = `
		INSERT INTO countries (countryISO2, countryName)
		VALUES (?, ?)
		ON DUPLICATE KEY UPDATE countryISO2 = VALUES(countryISO2);
	`

	CREATE_BANK_QUERY = `
		INSERT INTO banks (swiftCode, name, address, countryISO2, isHeadquarter, headquarterSwiftCode)
		VALUES (?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE swiftCode = VALUES(swiftCode);
	`

)
