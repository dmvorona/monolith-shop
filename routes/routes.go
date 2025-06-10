package routes

import (
	"github.com/gin-gonic/gin"
	"monolith-shop/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", controllers.Home)

	user := r.Group("/user")
	{
		user.POST("/register", controllers.RegisterUser)
	}

	product := r.Group("/products")
	{
		product.GET("/", controllers.ListProducts)
		product.POST("/", controllers.CreateProduct)
	}

	order := r.Group("/order")
	{
		order.POST("/", controllers.PlaceOrder)
	}

	cart := r.Group("/cart")
	{
		cart.POST("/add", controllers.AddToCart)
		cart.GET("/", controllers.ViewCart)
		cart.DELETE("/remove/:id", controllers.RemoveFromCart)
	}

}
