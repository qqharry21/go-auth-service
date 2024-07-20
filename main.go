package main

import (
	"go-auth-service/databases"
	"go-auth-service/middlewares"
	"go-auth-service/routes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/sirupsen/logrus"
)

// @title Go Auth Service
// @version 1.0
// @description "This is a simple authentication service"
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @host {{.Host}}
// @BasePath {{.BasePath}}
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error(err)
	}

	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(middlewares.NewCors([]string{"*"}))
	r.GET("swagger/*any", middlewares.NewSwagger())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	publicRoute := r.Group(os.Getenv("BASE_PATH"))
	resource, err := databases.InitResource()

	if err != nil {
		logrus.Error(err)
	}
	defer resource.Close()
	routes.InitUserRouter(publicRoute, resource)
	// routes.InitWorkflowRouter(publicRoute, resource)

	r.Run(":" + os.Getenv("PORT"))
}
