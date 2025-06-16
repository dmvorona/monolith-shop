package controllers

import (
	"github.com/gin-gonic/gin"
	"monolith-shop/db"
	"monolith-shop/models"
	"net/http"
)

func ListProducts(c *gin.Context) {
	var products []models.Product
	db.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&product)
	c.JSON(http.StatusOK, product)
}
