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
	DELETE_BANK_QUERY = `
		DELETE FROM banks WHERE swiftCode = ?;
	`

	GET_COUNTRY_QUERY = `
		SELECT * FROM countries WHERE countryISO2 = ?;
	`

	GET_BANKS_BY_COUNTRY_QUERY = `
		SELECT b.swiftCode, b.name, b.address, b.countryISO2, countries.countryName, b.isHeadquarter, b.headquarterSwiftCode
		FROM banks as b
		INNER JOIN countries ON b.countryISO2=countries.countryISO2
		WHERE b.countryISO2 = ?
	`
	GET_BANK_BY_SWIFTCODE_QUERY = `
		SELECT b.swiftCode, b.name, b.address, b.countryISO2, countries.countryName, b.isHeadquarter, b.headquarterSwiftCode
		FROM banks as b
		INNER JOIN countries ON b.countryISO2=countries.countryISO2
		WHERE swiftCode = ?
	`
	GET_BRANCHES_BY_HQ_SWIFTCODE_QUERY = `
		SELECT b.swiftCode, b.name, b.address, b.countryISO2, countries.countryName, b.isHeadquarter, b.headquarterSwiftCode
		FROM banks as b
		INNER JOIN countries ON b.countryISO2=countries.countryISO2
		WHERE headquarterSwiftCode = ?
	`
)
