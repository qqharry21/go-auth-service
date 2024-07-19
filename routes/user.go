package routes

import (
	"go-auth-service/databases"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(routerGroup *gin.RouterGroup, resource *databases.Resource) {
	userRouter := routerGroup.Group("/user")
	{
		userRouter.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "User",
			})
		})

		userRouter.POST("/register", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Register",
			})
		})

		userRouter.POST("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Login",
			})
		})
	}
}
