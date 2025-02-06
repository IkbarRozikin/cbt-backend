package main

import (
	"cbt-backend/database"
	"cbt-backend/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	app := fiber.New()

	database.Connect()

	app.Use(logger.New()) // Log setiap request
	app.Use(recover.New())

	// Inisialisasi handler
	authHandler := handlers.NewAuthHandler()
	//

	// Rute
	api := app.Group("/api/v1")
	authHandler.RegisterRoute(api.Group("/auth"))

	log.Fatal(app.Listen(":3000"))

}
