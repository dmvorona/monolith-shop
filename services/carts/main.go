package main

import (
	"github.com/dmvorona/shop/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/orders", controllers.ListOrders)
	r.POST("/orders", controllers.CreateOrder)

	r.Run(":8084")
}
