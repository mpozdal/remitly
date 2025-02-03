package config

import (
	"fmt"
	"os"
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

	return Config{
		PublicHost:  getEnv("PUBLIC_HOST", os.Getenv("PUBLIC_HOST")),
		Port:        getEnv("PORT", "8080"),
		DBName:      getEnv("DB_NAME", "swift"),
		DBAddress:   fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBUser:      getEnv("DB_USER", "root"),
		DBPassword:  getEnv("DB_PASSWORD", "mypassword"),
		CSVFilePath: getEnv("CSV_FILE_PATH", "assets/data.csv"),
	}
}
func getEnv(key, fallback string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback

}
