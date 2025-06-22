package main

import (
	"github.com/dmvorona/shop/controllers"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.GET("/", controllers.ListUsers)
		userGroup.POST("/", controllers.RegisterUser)
	}

	r.GET("/healthz", controllers.HealthCheck)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(":8082")
}
