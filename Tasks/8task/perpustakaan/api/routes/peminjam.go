package routes

import (
	"perpustakaan_app/api/handlers"
	"perpustakaan_app/pkg/peminjam"

	"github.com/gofiber/fiber/v2"
)

func PeminjamRoutes(api fiber.Router, peminjamService peminjam.PeminjamService) {
	api.Get("/peminjam", handlers.GetPeminjams(peminjamService))
	api.Post("/peminjam", handlers.CreatePeminjam(peminjamService))
	api.Put("/peminjam/:id", handlers.UpdatePeminjam(peminjamService))
	api.Delete("/peminjam/:id", handlers.DeletePeminjam(peminjamService))
}
