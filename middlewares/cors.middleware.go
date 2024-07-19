package middlewares

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

// NewCors return new gin handler fuc to handle CORS request
func NewCors(allowedOrigins []string) gin.HandlerFunc {
	return cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	})
}
