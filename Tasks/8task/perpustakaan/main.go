package main

import (
	"log"
	handler "perpustakaan_app/api/handlers"
	"perpustakaan_app/api/routes"
	"perpustakaan_app/pkg/buku"
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

	// Setup dependency injection
	repo := buku.NewBukuRepository(db)
	svc := buku.NewBukuService(repo)
	h := handler.NewBukuHandler(svc)

	// Init Fiber
	app := fiber.New()

	// Register routes
	api := app.Group("/api")
	routes.BukuRoutes(api, h)

	// Run server
	log.Fatal(app.Listen(":3000"))
}
