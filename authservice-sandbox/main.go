package main

import (
	"authservice-sandbox/api/adapter"

	"authservice-sandbox/api/routes"

	"authservice-sandbox/pkg/entities"

	"authservice-sandbox/pkg/user"

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
	db.AutoMigrate(&entities.User{})

	// TODO: 2. Inisialisasi repository, service, dan handler
	userRepo := user.NewUserRepository(db)
	jwtGenerator := adapter.NewJWTGenerator()
	userService := user.NewUserService(userRepo, jwtGenerator)

	// TODO: 3.Entry point aplikasi (misal: server HTTP)
	app := fiber.New()
	api := app.Group("/api")
	routes.UserRoutes(api, userService)

	// TODO: 4. Menjalankan server
	app.Listen(":8080")
}
