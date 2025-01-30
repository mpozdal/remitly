package main

import (
	"log"

	"mpozdal/remitly/cmd/api"
	"mpozdal/remitly/config"
	"mpozdal/remitly/db"

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
