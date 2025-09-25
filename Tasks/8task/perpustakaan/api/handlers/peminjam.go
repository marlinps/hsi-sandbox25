package handlers

import (
	"perpustakaan_app/pkg/entities"
	"perpustakaan_app/pkg/peminjam"

	"github.com/gofiber/fiber/v2"
)

func GetPeminjams(service peminjam.PeminjamService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		peminjams, err := service.FetchPeminjam()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to load peminjams",
			})
		}
		return c.JSON(peminjams)
	}
}

func CreatePeminjam(service peminjam.PeminjamService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var peminjam entities.Peminjam
		if err := c.BodyParser(&peminjam); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		createdPeminjam, err := service.InsertPeminjam(&peminjam)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create peminjam",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(createdPeminjam)
	}
}

func UpdatePeminjam(service peminjam.PeminjamService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var peminjam entities.Peminjam
		if err := c.BodyParser(&peminjam); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		updatedPeminjam, err := service.UpdatePeminjam(&peminjam)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update peminjam",
			})
		}
		return c.JSON(updatedPeminjam)
	}
}

func DeletePeminjam(service peminjam.PeminjamService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil || id <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid peminjam ID",
			})
		}
		err = service.RemovePeminjam(uint(id))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete peminjam",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Peminjam deleted successfully",
		})
	}
}
