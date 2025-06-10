package controllers

import (
	"github.com/gin-gonic/gin"
	"monolith-shop/db"
	"monolith-shop/models"
	"net/http"
)

func PlaceOrder(c *gin.Context) {
	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check product exists and has enough stock
	var product models.Product
	if err := db.DB.First(&product, input.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if product.Stock < input.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock"})
		return
	}

	// Create order
	order := models.Order{
		UserID:    input.UserID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
	}

	db.DB.Create(&order)

	// Update stock
	product.Stock -= input.Quantity
	db.DB.Save(&product)

	c.JSON(http.StatusOK, order)
}
