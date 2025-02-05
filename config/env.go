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

// 初始化配置
func initConfig() Config {
	// 加载.env文件
	err := godotenv.Load()
	// 如果加载失败，则输出错误信息并退出程序
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 返回配置信息
	return Config{
		// 公共主机
		PublicHost:  os.Getenv("PUBLIC_HOST"),
		Port:        os.Getenv("PORT"),
		DBName:      os.Getenv("DB_NAME"),
		DBAddress:   fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		CSVFilePath: os.Getenv("CSV_FILE_PATH"),
	}
}
