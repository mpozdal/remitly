package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost  string
	Port        string
	DBName      string
	DBAddress   string
	DBUser      string
	DBPassword  string
	CSVFilePath string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		PublicHost:  os.Getenv("PUBLIC_HOST"),
		Port:        os.Getenv("PORT"),
		DBName:      os.Getenv("DB_NAME"),
		DBAddress:   fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		CSVFilePath: os.Getenv("CSV_FILE_PATH"),
	}
}
