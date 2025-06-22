package controllers

import (
	"net/http"

	"github.com/dmvorona/shop/db"
	"github.com/dmvorona/shop/models"
	"github.com/gin-gonic/gin"
)

// GenerateReport provides a summary of all orders
func GenerateReport(c *gin.Context) {
	var orders []models.Order
	if err := db.DB.Preload("OrderItems.Product").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
