package main

import (
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	migrate "github.com/golang-migrate/migrate/v4"
	mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"mpozdal/remitly/config"
	"mpozdal/remitly/db"
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

	driver, err := mysqlMigrate.WithInstance(dbm.DB, &mysqlMigrate.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Użyj absolutnej ścieżki do migracji w kontenerze
	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/migrations", // Zmieniono na absolutną ścieżkę
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	v, d, _ := m.Version()
	log.Printf("Version: %d, dirty: %v", v, d)

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
