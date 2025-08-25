package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// TODO: inisialisasi framework Fiber
	app := fiber.New()

	// TODO: routing
	// URL http://localhost:3000/api/
	// Return "Ahlan Wa Sahlan"
	app.Get("/api/", func(c *fiber.Ctx) error {
		return c.SendString("Ahlan Wa Sahlan")
	})

	// Reserved PORT
	app.Listen(":3000")
}
