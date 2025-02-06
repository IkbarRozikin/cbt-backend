package handlers

import (
	"cbt-backend/database"
	"cbt-backend/models"
	"cbt-backend/repositories"
	"cbt-backend/services"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler() *AuthHandler {
	userRepo := repositories.NewUserRepository(database.DB)
	authService := services.NewAuthService(userRepo)

	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) RegisterRoute(router fiber.Router) {
	router.Post("/register", h.Register)
	router.Post("/login", h.Login)
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	ctx := c.Context()

	if err := h.authService.Register(ctx, &user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"code":    fiber.StatusCreated,
		"message": "User registered successfully",
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req := &models.Login{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Mendapatkan context dari Fiber dan meneruskannya
	ctx := c.Context()

	token, err := h.authService.Login(ctx, req.Username, req.Password)
	if err != nil {
		if err.Error() == "invalid password" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	return c.JSON(fiber.Map{
		"status":  "succes",
		"message": "login berhasil",
		"code":    fiber.StatusOK,
		"token":   token,
	})
}
