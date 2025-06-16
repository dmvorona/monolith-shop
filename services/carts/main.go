package main

import (
	"github.com/dmvorona/shop/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cart := r.Group("/carts")
	{
		cart.POST("/add", controllers.AddToCart)
		cart.GET("/", controllers.ViewCart)
		cart.DELETE("/remove/:id", controllers.RemoveFromCart)
	}

	r.Run(":8083")
}
