package main

import (
	"log"
	"os"

	pb "api-gateway/pb"
	"api-gateway/src/routes"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// Set up Gin router
	r := gin.Default()

	// Fetch HTTP port from environment variable or use default
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080" // Default HTTP port
	}

	// Run the HTTP server
	go func() {
		if err := r.Run(":" + httpPort); err != nil {
			log.Fatalf("Failed to run HTTP server: %v", err)
		}
	}()

	// Fetch gRPC host and port from environment variables or use defaults
	grpcAddr := os.Getenv("GRPC_ORDER_SERVICE")
	if grpcAddr == "" {
		grpcAddr = "localhost:50052" // Default gRPC address
	}

	// Connect to Oder Service via gRPC
	userConn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to user service at %s: %v", grpcAddr, err)
	}
	defer userConn.Close()
	userServiceClient := pb.NewUserServiceClient(userConn)

	// Initialize routes and pass the gRPC client
	routes.InitializeRoutes(r, userServiceClient)

	// Wait indefinitely
	select {}
}
