package handlers

import (
	"perpustakaan_app/pkg/buku"
	"perpustakaan_app/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

func AddBuku(service *buku.BukuService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var buku entities.Buku
		if err := c.BodyParser(&buku); err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		err := service.CreateBuku(buku)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create book",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(buku)
	}
}
