package main

import (
	"github.com/dmvorona/shop/controllers"
	"github.com/dmvorona/shop/db"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()
	db.InitDB()

	order := r.Group("/orders")
	{
		order.GET("/", controllers.ListOrders)
		order.POST("/", controllers.CreateOrder)
		order.GET("/healthz", controllers.HealthCheck)
	}

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(":8084")
}
