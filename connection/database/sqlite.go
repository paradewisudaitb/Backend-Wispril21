package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SQLiteConnect() *gorm.DB {
	if dbConnection == nil {
		dTemp, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
		dbConnection = dTemp
		if err != nil {
			panic(err)
		}
	}
	return dbConnection
}
