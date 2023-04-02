package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home.gohtml", gin.H{})
}

func LoadShopPage(c *gin.Context) {
	c.HTML(http.StatusOK, "shop.gohtml", gin.H{})
}

func LoadShopCartPage(c *gin.Context) {
	c.HTML(http.StatusOK, "shop-cart.gohtml", gin.H{})
}

func LoadAboutUsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about-us.gohtml", gin.H{})
}

func LoadContactUsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "contact-us.gohtml", gin.H{})
}
