package carts

import (
	"github.com/dmvorona/shop/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cart := r.Group("/cart")
	{
		cart.POST("/add", controllers.AddToCart)
		cart.GET("/view", controllers.ViewCart)
		cart.DELETE("/remove/:id", controllers.RemoveFromCart)
	}

	r.Run(":8080")
}
