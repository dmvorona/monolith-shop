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
		user.GET("/", controllers.ListUsers)
		user.POST("/report", controllers.GenerateReport)
		user.POST("/login", controllers.Login)
		user.POST("/logout", controllers.Logout)
	}

	product := r.Group("/products")
	{
		product.GET("", controllers.ListProducts)
		product.POST("", controllers.CreateProduct)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)
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
