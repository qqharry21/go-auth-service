package routes

import (
	"go-auth-service/controllers"
	"go-auth-service/databases"
	"go-auth-service/middlewares"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(routerGroup *gin.RouterGroup, resource *databases.Resource) {
	userController := controllers.NewUserController(resource)

	commonGroup := routerGroup.Group("")
	// commonGroup.POST("register", userController.Register)
	commonGroup.POST("login", userController.Login)
	// commonGroup.POST("refresh-token", userController.RefreshToken)

	authorizedGroup := routerGroup.Group("")
	authorizedGroup.Use(middlewares.JWTAuthMiddleware(resource.Redis))
	// authorizedGroup.POST("logout", userController.Logout)
}
