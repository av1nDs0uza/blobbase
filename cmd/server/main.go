package main

import (
	"log"

	"filestorage/internal/config"

	"github.com/gin-gonic/gin"

	"filestorage/internal/handler"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "running",
		})
	})
	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)

	log.Println("Server started on :8081")
	r.Run(":8081")

}
