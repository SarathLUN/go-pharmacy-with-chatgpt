package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDB() {
	dsn := "host=localhost user=root password=secret dbname=pharmacy_db port=5432 sslmode=disable TimeZone=Asia/Phnom_Penh"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Could not connect database: ", err.Error())
	}
	log.Println("DB connected: ", db)
}
