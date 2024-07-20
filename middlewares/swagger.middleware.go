package middlewares

import (
	"fmt"
	"go-auth-service/docs"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewSwagger() gin.HandlerFunc {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error(err)
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	basePath := os.Getenv("BASE_PATH")

	if host == "" || port == "" || basePath == "" {
		logrus.Error("Please set HOST, PORT, and BASE_PATH in .env")
		return nil
	}

	docs.SwaggerInfo.Title = "Go Auth Service"
	docs.SwaggerInfo.Description = "This is a simple authentication service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", host, port)
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}
