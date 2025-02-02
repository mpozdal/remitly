CREATE TABLE IF NOT EXISTS banks (
    swiftCode VARCHAR(11) PRIMARY KEY,
    name VARCHAR(255),
    address TEXT,
    countryISO2 VARCHAR(2),
	isHeadquarter BOOLEAN DEFAULT 1,
	headquarterSwiftCode VARCHAR(11) NULL,
    FOREIGN KEY (countryISO2) REFERENCES countries(countryISO2),
    FOREIGN KEY (headquarterSwiftCode) REFERENCES banks(swiftCode)

);