package handlers

import (
	"authservice-sandbox/pkg/user"

	"authservice-sandbox/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(service *user.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: 1. Menerima request dari user/client
		var req entities.User
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request",
			})
		}

		result, err := service.Login(req.Username, req.Password)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Pengguna tidak terdaftar",
			})
		}

		// TODO: 2. Mengembalikan response ke user/client
		return c.JSON(fiber.Map{
			"access_token": result,
		})
	}
}
