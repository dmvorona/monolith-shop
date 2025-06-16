package users

import (
	"github.com/dmvorona/shop/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.GET("/", controllers.ListUsers)
		userGroup.POST("/", controllers.RegisterUser)
	}

	r.Run(":8080")
}
