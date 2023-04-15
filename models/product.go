package models

import (
	"github.com/SarathLUN/go-pharmacy-with-chatgpt/config"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	IsFeatured  bool   `json:"is_featured"`
}

/*func init() {
	db := config.ConnectDB()
	err := db.AutoMigrate(&Product{})
	if err != nil {
		log.Fatalln("Unable to migrate the table product", err)
	}
}*/

func GetFeaturedProducts() []Product {
	db := config.ConnectDB()
	// Get all records
	var products []Product
	result := db.Find(&products)
	// SELECT * FROM users;
	if result.Error != nil {
		panic(result.Error)
	}

	return products
}
