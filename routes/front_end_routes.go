package routes

import (
	"github.com/SarathLUN/go-pharmacy-with-chatgpt/controllers"
	"github.com/gin-gonic/gin"
)

func LoadFrontEndRoutes(router *gin.Engine) {

	router.GET("/", controllers.LoadHomePage)
	router.GET("/shop", controllers.LoadShopPage)
	router.GET("/shop-cart", controllers.LoadShopCartPage)
	router.GET("/about-us", controllers.LoadAboutUsPage)
	router.GET("/contact-us", controllers.LoadContactUsPage)
}
