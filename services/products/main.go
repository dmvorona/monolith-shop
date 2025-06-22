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

	product := r.Group("/products")
	{
		product.GET("", controllers.ListProducts)
		product.POST("", controllers.CreateProduct)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)
		product.GET("/healthz", controllers.HealthCheck)
	}

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(":8081")
}
