package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	var router *gin.Engine = gin.Default()

	router.SetTrustedProxies(nil)
	// Define a simple GET endpoint for testing
	router.GET("/", func(c *gin.Context) {
		// Respond with a JSON message
		c.JSON(200, gin.H{
			"message": "Todo api is running!",
			"status":  "success",
		})
	})

	if err := router.Run(":8000"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
	log.Println("Server is running on http://localhost:8000")

}
