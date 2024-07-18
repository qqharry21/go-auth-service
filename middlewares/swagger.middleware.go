package middleware

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewSwagger() gin.HandlerFunc {
	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}
