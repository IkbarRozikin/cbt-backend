package handlers

import (
	"cbt-backend/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type userHandler struct {
	service services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{
		service: userService,
	}
}

func (h *userHandler) RegisterRoute(router fiber.Router) {
	router.Get("/:id", h.getUserById)
	router.Patch("/:id", h.updateUser)
	router.Delete("/:id", h.deleteUser)

}

func (h *userHandler) getUserById(c *fiber.Ctx) error {
	ctx := c.Context()

	// Ambil ID dari URL parameter
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"error":  "Invalid user ID",
		})
	}

	// Panggil service untuk mendapatkan user
	data, err := h.service.GetUserById(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": false,
			"error":  err.Error(),
		})
	}

	// Return response dengan status 200 (OK)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"code":   fiber.StatusOK,
		"data":   data,
	})
}

func (h *userHandler) updateUser(c *fiber.Ctx) error {
	// Ambil ID dari URL parameter
	idParam := c.Params("id")
	userID, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"error":  "Invalid user ID",
		})
	}

	// Decode request body ke map
	var input map[string]any
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	// Cek apakah ada data yang dikirim
	if len(input) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No fields to update"})
	}

	// Kirim ke service untuk update user
	if err := h.service.UpdateUser(c.Context(), userID, input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User updated successfully"})
}

func (h *userHandler) deleteUser(c *fiber.Ctx) error {
	// Ambil ID dari URL parameter
	idParam := c.Params("id")
	userID, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"error":  "Invalid user ID",
		})
	}

	if err := h.service.DeleteUser(c.Context(), userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User deleted successfully"})

}
