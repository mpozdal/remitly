package swift

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"mpozdal/remitly/db"
	"mpozdal/remitly/types"
)

var service *SwiftService

func initDatabase(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS countries (
			countryISO2 CHAR(2) PRIMARY KEY,
			countryName VARCHAR(100) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS banks (
			swiftCode VARCHAR(11) PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			countryISO2 CHAR(2) NOT NULL,
			address VARCHAR(255) NOT NULL,
			isHeadquarter BOOLEAN NOT NULL,
			headquarterSwiftCode VARCHAR(11),
			FOREIGN KEY (countryISO2) REFERENCES countries(countryISO2)
		);`,
		`INSERT INTO countries (countryISO2, countryName) VALUES ('US', 'United States') ON DUPLICATE KEY UPDATE countryName=VALUES(countryName);`,
		`INSERT INTO banks (swiftCode, name, countryISO2, address, isHeadquarter) VALUES ('BOFAUS3N', 'Bank of America', 'US', 'Street 1', true) ON DUPLICATE KEY UPDATE name=VALUES(name);`,
	}

	for _, query := range queries {

		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to execute query: %v %s", err, query)
		}
	}

	return nil
}

func setupTestContainer(t *testing.T) func() {
	dsn := "root:password@tcp(db_test:3306)/testdb?parseTime=true"

	testDB, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to test database")

	dbManager, err := db.NewDBManagerWithCon(testDB)
	if err != nil {
		t.Fatalf("Failed to create DBManager: %v", err)
	}

	err = initDatabase(testDB)
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	service = NewSwiftService(dbManager)

	return func() {
		teardown()
		_ = testDB.Close()
	}
}
func teardown() {
	cmd := exec.Command("docker", "stop", "remitly-db_test-1")
	cmd.Run()

}

func TestAddNewData(t *testing.T) {
	teardown := setupTestContainer(t)
	defer teardown()

	payload := types.AddBankPayload{
		SwiftCode:     "TEST1234",
		BankName:      "Test Bank",
		CountryISO2:   "US",
		CountryName:   "United States",
		IsHeadquarter: true,
	}

	response := service.AddNewData(payload)
	assert.Equal(t, 200, response.Status)
}

func TestGetDataBySwiftCode2(t *testing.T) {
	teardown := setupTestContainer(t)
	defer teardown()

	payload := types.AddBankPayload{
		SwiftCode:     "TEST1234",
		BankName:      "Test Bank",
		CountryISO2:   "US",
		CountryName:   "United States",
		IsHeadquarter: true,
	}
	service.AddNewData(payload)

	bank, err := service.GetDataBySwiftCode("TEST1234")
	assert.NoError(t, err)
	assert.Equal(t, "Test Bank", bank.BankName)
}

func TestDeleteData(t *testing.T) {
	teardown := setupTestContainer(t)
	defer teardown()

	payload := types.AddBankPayload{
		SwiftCode:     "TEST1234",
		BankName:      "Test Bank",
		CountryISO2:   "US",
		CountryName:   "United States",
		IsHeadquarter: true,
	}
	service.AddNewData(payload)

	response := service.DeleteData("TEST1234")
	assert.Equal(t, 200, response.Status)

}
