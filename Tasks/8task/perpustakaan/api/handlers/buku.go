package handlers

import (
	"perpustakaan_app/pkg/buku"

	"perpustakaan_app/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

func GetBukus(service buku.BukuService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bukus, err := service.FetchBuku()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to load books",
			})
		}
		return c.JSON(bukus)
	}
}

func CreateBuku(service buku.BukuService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var buku entities.Buku
		if err := c.BodyParser(&buku); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		createdBuku, err := service.InsertBuku(&buku)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create book",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(createdBuku)
	}
}

func UpdateBuku(service buku.BukuService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var buku entities.Buku
		if err := c.BodyParser(&buku); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		updatedBuku, err := service.UpdateBuku(&buku)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update book",
			})
		}
		return c.JSON(updatedBuku)
	}
}

func DeleteBuku(service buku.BukuService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil || id <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid book ID",
			})
		}
		err = service.RemoveBuku(uint(id))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete book",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Book deleted successfully",
		})
	}
}
