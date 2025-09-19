package routes

import (
	"book_rental_app/pkg/buku"

	"book_rental_app/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func BukuRoutes(api fiber.Router, bukuService *buku.BukuService) {
	api.Get("/bukus", handlers.LoadAllBuku(bukuService))
	api.Post("/bukus", handlers.BuatBukuBaru(bukuService))
}
