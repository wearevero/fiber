package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/wearevero/fiber/models"
	"github.com/wearevero/fiber/routes"
)

func main() {
	// Connect to database
	models.ConnectDatabase()

	// get value from APP_PORT
	port := os.Getenv("APP_PORT")
	if port == "" {
		// set default port to 8000
		port = "8000"
	}

	// Create instance Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Run server with defined port
	log.Printf("Server running on port %s", port)
	err := app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
