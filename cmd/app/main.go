package main

import (
	"fastx-api/config"
	"fastx-api/src/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Set up Gin router
	r := gin.Default()

	// Initialize routes
	routes.InitializeRoutes(r)

	// Run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
