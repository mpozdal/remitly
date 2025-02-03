package main

import (
	"bufio"
	"fmt"
	"log"
	"mpozdal/remitly/cmd/api"
	"mpozdal/remitly/config"
	"mpozdal/remitly/db"
	"mpozdal/remitly/services/csvparser"
	"mpozdal/remitly/types"
	"mpozdal/remitly/utils"
	"os"

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
	log.Println()
	if err != nil {
		log.Fatal(err)
	}
	initDatabase(dbm)
	for {
		fmt.Println("Do you want to load data from file? (y/n)")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input := scanner.Text()
			if input == "Y" || input == "y" {
				csvparser := csvparser.NewCSVParser(config.Envs.CSVFilePath)

				countries, banks, err := csvparser.ParseRecords()

				if err != nil {
					log.Fatal("Error reading CSV file", err)

				}
				err = insertData(dbm, countries, banks)
				if err != nil {
					log.Println("Error inserting data", err)
				}
				fmt.Println("Data has been loaded")
				break
			} else if input == "N" || input == "n" {
				fmt.Println("Load data has been cancelled")
				break
			} else {
				fmt.Println("Input 'Y' or 'N'")

			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	server := api.NewAPIServer(":"+config.Envs.Port, dbm)
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
