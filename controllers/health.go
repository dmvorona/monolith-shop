package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
