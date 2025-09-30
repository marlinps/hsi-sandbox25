package main

import (
	"log"

	"perpustakaan_app/api/routes"

	"perpustakaan_app/pkg/buku"

	"perpustakaan_app/pkg/peminjam"

	"perpustakaan_app/pkg/entities"

	"github.com/gofiber/fiber/v2"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func main() {
	// TODO: 1. Konfigurasi Database dan inisialisasi repository serta service
	dsn := "root:@tcp(127.0.0.1:3306)/sandbox_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal Terhubung ke Database")
	}

	// Auto migrate entity
	db.AutoMigrate(&entities.Buku{})
	db.AutoMigrate(&entities.Peminjam{})

	// Setup dependency injection
	bukuRepository := buku.NewBukuRepository(db)
	peminjamRepository := peminjam.NewPeminjamRepository(db)

	bukuService := buku.NewBukuService(bukuRepository)
	peminjamService := peminjam.NewPeminjamService(peminjamRepository)

	// Init Fiber
	app := fiber.New()

	// Register routes
	api := app.Group("/api")
	routes.BukuRoutes(api, bukuService)
	routes.PeminjamRoutes(api, peminjamService)

	// Run server
	log.Fatal(app.Listen(":3000"))
}
