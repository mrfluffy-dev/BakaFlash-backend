package main

import (
	"BakaFlash/api" // Import the "api" package
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize API routes by calling SetupRoutes
	api.SetupRoutes(r)

	// Start the server
	r.Run(":8080")
}
