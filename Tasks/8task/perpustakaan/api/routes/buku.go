package routes

import (
	"perpustakaan_app/pkg/buku"

	"perpustakaan_app/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func BukuRoutes(api fiber.Router, bukuService buku.BukuService) {
	api.Get("/buku", handlers.GetBukus(bukuService))
	api.Post("/buku", handlers.CreateBuku(bukuService))
	api.Put("/buku/:id", handlers.UpdateBuku(bukuService))
	api.Delete("/buku/:id", handlers.DeleteBuku(bukuService))
}
