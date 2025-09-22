package handler

import (
	"net/http"
	"strconv"

	"perpustakaan_app/pkg/entities"

	"perpustakaan_app/pkg/buku/service"

	"github.com/gofiber/fiber/v2"
)

type BukuHandler struct {
	service service.BukuService
}

// Constructor
func NewBukuHandler(service service.BukuService) *BukuHandler {
	return &BukuHandler{service: service}
}

func (h *BukuHandler) CreateBuku(c *fiber.Ctx) error {
	var buku entities.Buku
	if err := c.BodyParser(&buku); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	created, err := h.service.CreateBuku(&buku)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(created)
}

func (h *BukuHandler) GetAllBuku(c *fiber.Ctx) error {
	bukus, err := h.service.GetAllBuku()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(bukus)
}

func (h *BukuHandler) UpdateBuku(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	var buku entities.Buku
	if err := c.BodyParser(&buku); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	buku.ID = uint(id)

	updated, err := h.service.UpdateBuku(&buku)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updated)
}

func (h *BukuHandler) DeleteBuku(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	if err := h.service.DeleteBuku(uint(id)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(http.StatusNoContent)
}
