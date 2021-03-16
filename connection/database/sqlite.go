package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SQLiteConnect() *gorm.DB {
	if database == nil {
		dTemp, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
		database = dTemp
		if err != nil {
			panic(err)
		}
	}
	return database
}
