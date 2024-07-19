package middlewares

import (
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

	docs.SwaggerInfo.Title = "Go Auth Service"
	docs.SwaggerInfo.Description = "This is a simple authentication service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("HOST") + ":" + os.Getenv("PORT")
	docs.SwaggerInfo.BasePath = os.Getenv("BASE_PATH")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}
