package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresConnect is function to make connection
func PostgresConnect() *gorm.DB {
	if database == nil {
		godotenv.Load()
		host := os.Getenv("PG_HOST")
		port := os.Getenv("PG_PORT")
		dbname := os.Getenv("PG_DATABASE")
		user := os.Getenv("PG_USERNAME")
		password := os.Getenv("PG_PASSWORD")
		psqlLoginInfo := fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
			host, port, user, password, dbname)
		dTemp, err := gorm.Open(postgres.Open(psqlLoginInfo), &gorm.Config{})
		database = dTemp
		if err != nil {
			panic(err)
		}
	}
	return database
}
