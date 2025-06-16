package products

import (
	"github.com/dmvorona/shop/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/products", controllers.ListProducts)
	r.POST("/products", controllers.CreateProduct)

	r.Run(":8080")
}
