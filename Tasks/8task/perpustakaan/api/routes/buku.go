package routes

import (
	"perpustakaan_app/pkg/buku"

	"perpustakaan_app/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func BukuRoutes(api fiber.Router, bukuService *buku.BukuService) {
	api.Get("/bukus", handlers.GetBukus(bukuService))
	api.Post("/bukus", handlers.CreateBuku(bukuService))
	api.Put("/bukus/:id", handlers.UpdateBuku(bukuService))
	api.Delete("/bukus/:id", handlers.DeleteBuku(bukuService))
}
