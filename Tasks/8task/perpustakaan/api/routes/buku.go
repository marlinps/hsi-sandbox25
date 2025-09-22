package routes

import (
	handler "perpustakaan_app/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func BukuRoutes(router fiber.Router, h *handler.BukuHandler) {
	api := router.Group("/buku")

	api.Post("/", h.CreateBuku)
	api.Get("/", h.GetAllBuku)
	api.Put("/:id", h.UpdateBuku)
	api.Delete("/:id", h.DeleteBuku)
}
