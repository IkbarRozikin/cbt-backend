package main

import (
	"cbt-backend/db"
	"cbt-backend/handlers"
	"cbt-backend/repositories"
	"cbt-backend/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Inisialisasi Fiber
	app := fiber.New()

	// Koneksi ke Database
	db.Connect()

	// Middleware
	app.Use(logger.New())  // Logging request
	app.Use(recover.New()) // Menangani panic agar tidak crash

	// Inisialisasi Repository
	userRepo := repositories.NewUserRepository(db.DB)

	// Inisialisasi Service
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)

	// Inisialisasi Handler
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	// Rute API
	api := app.Group("/api/v1")
	authHandler.RegisterRoute(api.Group("/auth"))
	userHandler.RegisterRoute(api.Group("/users"))

	// Jalankan Server
	log.Fatal(app.Listen(":3000"))
}
