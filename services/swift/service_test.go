package swift

import (
	"database/sql"
	"testing"

	"mpozdal/remitly/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetDataBySwiftCode(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	dbManager := &db.DBManager{DB: mockDB}
	service := NewSwiftService(dbManager)

	t.Run("Bank is headquarter", func(t *testing.T) {

		rows := sqlmock.NewRows([]string{"swift_code", "bank_name", "country_iso2", "country_name", "address", "is_headquarter", "headquarterSwiftCode"}).
			AddRow("ABCDEF", "Test Bank", "US", "United States", "123 Main St", true, sql.NullString{Valid: false})
		mock.ExpectQuery(db.GET_BANK_BY_SWIFTCODE_QUERY).WithArgs("ABCDEF").WillReturnRows(rows)

		branchRows := sqlmock.NewRows([]string{"swift_code", "bank_name", "country_iso2", "country_name", "address", "is_headquarter", "headquarterSwiftCode"}).
			AddRow("GHIJKL", "Test Branch", "US", "United States", "456 Branch St", false, sql.NullString{String: "ABCDEF", Valid: true})
		mock.ExpectQuery(db.GET_BRANCHES_BY_HQ_SWIFTCODE_QUERY).WithArgs("ABCDEF").WillReturnRows(branchRows)

		response, err := service.GetDataBySwiftCode("ABCDEF")

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, "ABCDEF", response.SwiftCode)
		assert.True(t, response.IsHeadquarter)
		assert.Len(t, response.Branches, 1)
	})

	t.Run("Bank is not headquarter", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"swift_code", "bank_name", "country_iso2", "country_name", "address", "is_headquarter", "headquarterSwiftCode"}).
			AddRow("ABCDEF", "Test Bank", "US", "United States", "123 Main St", false, sql.NullString{Valid: false})
		mock.ExpectQuery(db.GET_BANK_BY_SWIFTCODE_QUERY).WithArgs("ABCDEF").WillReturnRows(rows)

		response, err := service.GetDataBySwiftCode("ABCDEF")

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, "ABCDEF", response.SwiftCode)
		assert.False(t, response.IsHeadquarter)
		assert.Nil(t, response.Branches)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
