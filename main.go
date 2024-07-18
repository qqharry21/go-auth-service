package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	middleware "github.com/qqharry21/go-auth-service/middlewares"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var redisClient *redis.Client

func main() {
	// 連接到 MongoDB
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(context.Background())

	// 連接到 Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("swagger/*any", middleware.NewSwagger())

	r.Run(":8080")
}
