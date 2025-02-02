package main

import (
	"log"
	"mpozdal/remitly/cmd/api"
	"mpozdal/remitly/config"
	"mpozdal/remitly/db"
	"mpozdal/remitly/services/csvparser"
	"mpozdal/remitly/types"
	"mpozdal/remitly/utils"

	"github.com/go-sql-driver/mysql"
)

func main() {

	dbm, err := db.NewDBManager(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	initDatabase(dbm)

	csvparser := csvparser.NewCSVParser(config.Envs.CSVFilePath)

	countries, banks, err := csvparser.ParseRecords()

	if err != nil {
		log.Fatal("Error reading CSV file", err)

	}
	err = insertData(dbm, countries, banks)
	if err != nil {
		log.Println("Error inserting data", err)
	}

	server := api.NewAPIServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initDatabase(dbm *db.DBManager) {
	err := dbm.DB.Ping()
	if err != nil {
		log.Fatal(err)

	}

	log.Println("Database setup completed")

}

func insertData(dbm *db.DBManager, countries []types.Country, banks []types.Bank) error {

	for _, country := range countries {

		err := dbm.AddCountry(country)
		if err != nil {
			log.Fatal("Error inserting country", err)
		}

	}

	for _, bank := range utils.SortBanks(banks) {

		_, err := dbm.AddBank(bank)
		if err != nil {
			log.Fatal("Error inserting bank", err)
		}

	}

	return nil
}
