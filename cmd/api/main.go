package main

import (
	"fmt"
	"log"
	"todo-api/internal/config"
	"todo-api/internal/database"
	"todo-api/internal/handelers"
	"todo-api/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load()
	var cfg *config.Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DATABASE_URL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()
	// Initialize Gin router
	var router *gin.Engine = gin.Default()

	router.SetTrustedProxies(nil)
	// Define a simple GET endpoint for testing
	router.GET("/", func(c *gin.Context) {
		// Respond with a JSON message
		c.JSON(200, gin.H{
			"message":  "Todo api is running!",
			"status":   "success",
			"database": "connected",
		})
	})

	router.POST("/todos", handelers.NewTodoHandler(repository.NewTodoRepository(pool)).CreateTodo)

	if err := router.Run(":8000"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
	log.Println("Server is running on http://localhost:8000")

}
