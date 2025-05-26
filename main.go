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
	"github.com/wearevero/fiber/controllers/MasterData/usercontroller"
	"github.com/wearevero/fiber/models"
	"github.com/wearevero/fiber/routes"
)

func main() {
	// Connect to database
	models.ConnectDatabase()
	bagiancontroller.SetDB(models.DB)
	jabatancontroller.SetDB(models.DB)
	karyawancontroller.SetDB(models.DB)
	usercontroller.SetDB(models.DB)
	absenjamcontroller.SetDB(models.DB)
	absenhariancontroller.SetDB(models.DB)
	absenlemburcontroller.SetDB(models.DB)

	// Get value from APP_PORT
	port := os.Getenv("APP_PORT")
	if port == "" {
		// set default port to 8000
		port = "8000"
	}

	// Create instance Fiber app
	app := fiber.New()

	// CORS middleware
	app.Use(cors.New())

	// Does offer advanced configuration:
	/*
		app.Use(cors.New(cors.Config{
			AllowOrigins:     "http://localhost:3000",
			AllowMethods:     "GET, POST, PUT, DELETE",
			AllowHeaders:     "Origin, Content-Type, Accept",
			ExposeHeaders:    "Content-Length",
			AllowCredentials: true,
		}))
	*/

	// Setup routes
	routes.SetupRoutes(app)

	// Run server with defined port
	log.Printf("Server running on port %s", port)
	err := app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
