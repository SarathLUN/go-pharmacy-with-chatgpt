package main

import (
	"github.com/SarathLUN/go-pharmacy-with-chatgpt/config"
	"github.com/SarathLUN/go-pharmacy-with-chatgpt/models"
	"github.com/SarathLUN/go-pharmacy-with-chatgpt/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	// test DB connection
	db := config.ConnectDB()
	err := db.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalln("Error migrating database: ", err)
	}
}

func main() {
	// Set Gin to production mode
	//gin.SetMode(gin.ReleaseMode)

	// Create a new Gin router
	router := gin.Default()

	// Serve static files from the "static" directory
	router.Static("/static", "./static")
	router.LoadHTMLGlob("views/**/*")

	// Load front-end routes
	routes.LoadFrontEndRoutes(router)

	// Start the server
	if err := router.Run("localhost:8080"); err != nil {
		panic(err)
	}

}
