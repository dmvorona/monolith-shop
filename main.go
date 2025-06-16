package main

import (
	"github.com/dmvorona/monolith-shop/db"
	"github.com/dmvorona/monolith-shop/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	db.InitDB()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
