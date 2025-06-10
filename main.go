package main

import (
	"github.com/gin-gonic/gin"
	"monolith-shop/db"
	"monolith-shop/routes"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	db.InitDB()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
