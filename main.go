package main

import (
	"github.com/dmvorona/shop/db"
	"github.com/dmvorona/shop/routes"
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
