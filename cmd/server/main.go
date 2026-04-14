package main

import (
	"log"

	"filestorage/internal/config"
	"filestorage/internal/handler"
	"filestorage/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	r := gin.Default()

	// public routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "running"})
	})

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	// protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/profile", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"user_id": c.GetInt("user_id"),
			"email":   c.GetString("email"),
		})
	})

	log.Println("Server started on :8081")
	r.Run(":8081")
}
