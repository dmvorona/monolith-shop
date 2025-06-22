package main

import (
	"github.com/dmvorona/shop/controllers"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()

	r.GET("", controllers.ListProducts)
	r.POST("", controllers.CreateProduct)
	r.PUT("/:id", controllers.UpdateProduct)
	r.DELETE("/:id", controllers.DeleteProduct)

	r.GET("/healthz", controllers.HealthCheck)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(":8081")
}
