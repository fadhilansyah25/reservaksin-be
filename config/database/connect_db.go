package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

// postgresSQL
func (dbConfig *DBConfig) DbURLMain() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
	)
}

func Connection() {
	errorEnv := godotenv.Load(".env")
	if errorEnv != nil {
		log.Fatalf("Error loading .env file")
	}

	var dsn = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	DB, err = gorm.Open(postgres.Open(dsn.DbURLMain()), &gorm.Config{})

	if err != nil {
		fmt.Println("Status:", err)
	}
	Migrate()
}
