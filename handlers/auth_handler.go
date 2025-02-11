package handlers

import (
	"cbt-backend/database"
	"cbt-backend/models"
	"cbt-backend/repositories"
	"cbt-backend/services"
	"cbt-backend/validators"
	"fmt"

	"github.com/go-playground/validator/v10"
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
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	validators.Validate.RegisterValidation("usernameRegexp", validators.UsernameRegexp)

	// Validasi user struct
	err := validators.Validate.Struct(user)
	if err != nil {
		// Membuat map untuk menampung error
		validationErrors := make(map[string]string)
		// Iterasi error dan tampilkan penjelasan lengkap
		for _, err := range err.(validator.ValidationErrors) {
			// Menambahkan informasi lengkap tentang error
			validationErrors[err.Field()] = fmt.Sprintf("Invalid value for '%s', expected %s", err.Field(), err.Tag())
		}
		// Mengembalikan error dalam format JSON
		return c.Status(400).JSON(fiber.Map{
			"validation_errors": validationErrors,
		})
	}

	ctx := c.Context()

	if err := h.authService.Register(ctx, &user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status((fiber.StatusCreated)).JSON(fiber.Map{
		"status":  true,
		"code":    fiber.StatusCreated,
		"message": "User registered successfully",
	})
}
