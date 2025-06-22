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

	cart := r.Group("/carts")
	{
		cart.POST("/add", controllers.AddToCart)
		cart.GET("/", controllers.ViewCart)
		cart.DELETE("/remove/:id", controllers.RemoveFromCart)
		cart.GET("/healthz", controllers.HealthCheck)
	}

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(":8083")
}
