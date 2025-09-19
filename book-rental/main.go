package main

import (
	"book_rental_app/api/routes"
	"book_rental_app/pkg/buku"
	"book_rental_app/pkg/entities"

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
	db.AutoMigrate(&entities.Buku{})

	// TODO: 2. Inisialisasi Repository
	bukuRepository := buku.NewBukuRepository(db)

	// TODO: 3. Inisialisasi Service
	bukuService := buku.NewBukuService(bukuRepository)

	// TODO: 4. Inisialisasi Fiber dan Setup Routes
	app := fiber.New()

	api := app.Group("/api")
	routes.BukuRoutes(api, bukuService)
	app.Listen(":3000")
}
