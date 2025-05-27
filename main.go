package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/wearevero/fiber/controllers/Laporan/absenhariancontroller"
	"github.com/wearevero/fiber/controllers/Laporan/absenjamcontroller"
	"github.com/wearevero/fiber/controllers/Laporan/absenlemburcontroller"
	"github.com/wearevero/fiber/controllers/MasterData/bagiancontroller"
	"github.com/wearevero/fiber/controllers/MasterData/jabatancontroller"
	"github.com/wearevero/fiber/controllers/MasterData/karyawancontroller"
	"github.com/wearevero/fiber/controllers/MasterData/karyawankeluarcontroller"
	"github.com/wearevero/fiber/controllers/MasterData/usercontroller"
	"github.com/wearevero/fiber/models"
	"github.com/wearevero/fiber/routes"
)

func main() {
	// Connect to database
	models.ConnectDatabase()

	// Set DB instance to each controller (if required)
	db := models.DB
	bagiancontroller.SetDB(db)
	jabatancontroller.SetDB(db)
	karyawancontroller.SetDB(db)
	karyawankeluarcontroller.SetDB(db)
	usercontroller.SetDB(db)
	absenjamcontroller.SetDB(db)
	absenhariancontroller.SetDB(db)
	absenlemburcontroller.SetDB(db)

	// Get port from .env or use default
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	// Create new Fiber app
	app := fiber.New()

	// Setup CORS
	app.Use(cors.New())

	// Register API routes
	routes.RegisterAPIRoutes(app)

	// Start server
	log.Printf("Server running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
