package routes

import (
	"github.com/dmvorona/shop/controllers"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", controllers.Home)

	user := r.Group("/users")
	{
		user.POST("/register", controllers.RegisterUser)
	}

	product := r.Group("/products")
	{
		product.GET("/", controllers.ListProducts)
		product.POST("/", controllers.CreateProduct)
	}

	order := r.Group("/orders")
	{
		order.GET("/", controllers.ListOrders)
		order.POST("/", controllers.CreateOrder)
	}

	cart := r.Group("/carts")
	{
		cart.POST("/add", controllers.AddToCart)
		cart.GET("/", controllers.ViewCart)
		cart.DELETE("/remove/:id", controllers.RemoveFromCart)
	}

	r.GET("/healthz", controllers.HealthCheck)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
