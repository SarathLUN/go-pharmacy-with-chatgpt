package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadHomePage(c *gin.Context) {

	c.HTML(http.StatusOK, "home.gohtml", gin.H{
		"Active": "home",
	})
}

func LoadShopPage(c *gin.Context) {
	c.HTML(http.StatusOK, "shop.gohtml", gin.H{
		"Active": "shop",
	})
}

func LoadShopCartPage(c *gin.Context) {
	c.HTML(http.StatusOK, "shop-cart.gohtml", gin.H{
		"Active": "shop-cart",
	})
}

func LoadAboutUsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about-us.gohtml", gin.H{
		"Active": "about-us",
	})
}

func LoadContactUsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "contact-us.gohtml", gin.H{
		"Active": "contact-us",
	})
}
