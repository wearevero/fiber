package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/wearevero/fiber/models"
	"github.com/wearevero/fiber/routes"
)

func main() {
	// Menghubungkan ke database
	models.ConnectDatabase()

	// Mengambil nilai port dari environment variable
	port := os.Getenv("APP_PORT")
	if port == "" {
		// Jika tidak ada, set default ke port 8080
		port = "8080"
	}

	// Membuat instance aplikasi Fiber
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Menjalankan aplikasi di port yang ditentukan
	log.Printf("Server running on port %s", port)
	err := app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
