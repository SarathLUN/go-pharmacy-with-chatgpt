package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("pharmacy.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("DB connected: ", db)
	return db
}
