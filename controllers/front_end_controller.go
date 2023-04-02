package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func LoadShopPage(c *gin.Context) {
	c.HTML(http.StatusOK, "shop.html", gin.H{})
}

func LoadShopCartPage(c *gin.Context) {
	c.HTML(http.StatusOK, "shop-cart.html", gin.H{})
}

func LoadAboutUsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about-us.html", gin.H{})
}

func LoadContactUsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "contact-us.html", gin.H{})
}
