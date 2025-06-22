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

	user := r.Group("/users")
	{
		user.GET("/", controllers.ListUsers)
		user.POST("/register", controllers.RegisterUser)
		user.POST("/report", controllers.GenerateReport)
		user.POST("/login", controllers.Login)
		user.POST("/logout", controllers.Logout)
		user.GET("/healthz", controllers.HealthCheck)
	}

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(":8082")
}
