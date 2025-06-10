package controllers

import (
	"github.com/gin-gonic/gin"
	"monolith-shop/db"
	"monolith-shop/models"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}
