package main

import (
	"github.com/dmvorona/shop/controllers"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()

	order := r.Group("/orders")
	{
		order.GET("/", controllers.ListOrders)
		order.POST("/", controllers.CreateOrder)
	}

	r.GET("/healthz", controllers.HealthCheck)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(":8084")
}
