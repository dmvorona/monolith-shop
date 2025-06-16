package controllers

import (
	"github.com/dmvorona/shop/db"
	"github.com/dmvorona/shop/models"
	"github.com/gin-gonic/gin"
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

func ListUsers(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}
