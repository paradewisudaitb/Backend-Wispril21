package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

const driver = "postgres"

var database *gorm.DB

// PostgresConnect is function to make connection
func PostgresConnect() *gorm.DB {
	if database == nil {
		godotenv.Load()
		host := os.Getenv("PG_HOST")
		port := os.Getenv("PG_PORT")
		dbname := os.Getenv("PG_NAME")
		user := os.Getenv("PG_USER")
		password := os.Getenv("PG_PASSWORD")
		psqlLoginInfo := fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		dTemp, err := gorm.Open(driver, psqlLoginInfo)
		database = dTemp
		if err != nil {
			panic(err)
		}
	}
	return database
}
